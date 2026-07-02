package goroutines

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
)

type strErr string

func (s strErr) Error() string {
	return string(s)
}

const (
	TimeoutError     = strErr("timed out")
	QueueFullError   = strErr("queue full")
	PoolStoppedError = strErr("pool shutdown")
)

const (
	// shutdownPollInterval is how often ShutdownAndWait re-checks that the pool
	// has drained. It's coarse because shutdown is one-time and latency-tolerant.
	shutdownPollInterval = 50 * time.Millisecond

	// idlePollInterval is how often AwaitIdle re-checks for outstanding work.
	// It's short so callers that submit work and immediately wait stay responsive.
	idlePollInterval = 5 * time.Millisecond
)

// Pool represents a goroutine worker pool that can be configured with a queue size and min and max sizes.
//
//	The pool will start with min size goroutines and will add more if the queue isn't staying empty.
//	After a worker has been idle for a configured time, it will stop
type Pool interface {
	// Queue submits a unit of work to the pool. It will return an error if the pool is shutdown
	Queue(func()) error

	// QueueWithTimeout submits a unit of work to the pool. It will return an error if the pool is shutdown or
	// if the work cannot be submitted to the work queue before the given timeout elapses
	QueueWithTimeout(func(), time.Duration) error

	// QueueOrError submits a unit of work to the pool. It will return an error if the pool is shutdown or
	// if the work cannot be submitted to the work queue immediately
	QueueOrError(func()) error

	// GetWorkerCount returns the current number of goroutines servicing the work queue
	GetWorkerCount() uint32

	// GetQueueSize returns the current number of work items in the work queue
	GetQueueSize() uint32

	// GetBusyWorkers returns the current number workers busy doing work from the work queue
	GetBusyWorkers() uint32

	// GetOutstanding returns the number of work items that have been accepted but not yet completed,
	// that is, items still queued plus items currently running
	GetOutstanding() uint32

	// Shutdown stops all workers as they finish work and prevents new work from being submitted to the queue
	Shutdown()

	// ShutdownAndWait shuts the pool down (see Shutdown) and then blocks until all in-flight work has
	// completed and all workers have exited, or until the timeout elapses. Queued work that has not
	// yet started is abandoned, not run. Returns TimeoutError if the timeout elapses before the pool
	// drains, in which case some work may still be running.
	ShutdownAndWait(timeout time.Duration) error

	// AwaitIdle blocks until the pool has no outstanding work (nothing queued and nothing running),
	// or until the timeout elapses, returning TimeoutError on timeout. It is a point-in-time wait:
	// new work may be submitted concurrently, and if the pool has been shut down with work still
	// queued, that work is abandoned and AwaitIdle will time out.
	AwaitIdle(timeout time.Duration) error
}

// PoolConfig is used to configure a new Pool
type PoolConfig struct {
	// The size of the channel feeding the worker pool
	QueueSize uint32
	// The minimum number of goroutines
	MinWorkers uint32
	// The maximum number of workers
	MaxWorkers uint32
	// IdleTime how long a goroutine should be idle before exiting
	IdleTime time.Duration
	// Provides a way to join shutdown of the pool with other components.
	// The pool also be shut down independently using the Shutdown method
	CloseNotify <-chan struct{}
	// Provides a way to specify what happens if a worker encounters a panic
	// if no PanicHandler is provided, panics will not be caught
	PanicHandler func(err interface{})
	// Optional callback which is called whenever work completes, with the
	// time the work took to complete
	OnWorkCallback func(workTime time.Duration)
	// Optional callback which is called when the pool is created
	OnCreate func(Pool)
	// A function to identify the pool type. WorkerFunction must call embedded function to start the worker
	WorkerFunction func(uint32, func())
}

func (self *PoolConfig) Validate() error {
	if self.MaxWorkers < 1 {
		return fmt.Errorf("max workers must be at least 1")
	}
	if self.MinWorkers > self.MaxWorkers {
		return fmt.Errorf("min workers must be less than or equal to max workers. min workers=%v, max workers=%v", self.MinWorkers, self.MaxWorkers)
	}
	return nil
}

func NewPool(config PoolConfig) (Pool, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	if config.MinWorkers > math.MaxInt32 {
		return nil, fmt.Errorf("min workers must be less than or equal to %v", math.MaxInt32)
	}

	if config.MaxWorkers > math.MaxInt32 {
		return nil, fmt.Errorf("max workers must be less than or equal to %v", math.MaxInt32)
	}

	result := &pool{
		queue:               make(chan func(), int(config.QueueSize)),
		minWorkers:          int32(config.MinWorkers),
		maxWorkers:          int32(config.MaxWorkers),
		maxIdle:             config.IdleTime,
		externalCloseNotify: config.CloseNotify,
		closeNotify:         make(chan struct{}),
		panicHandler:        config.PanicHandler,
		onWorkCallback:      config.OnWorkCallback,
		workF:               config.WorkerFunction,
	}

	if result.workF == nil {
		result.workF = func(workerNumber uint32, worker func()) {
			worker()
		}
	}

	if config.OnCreate != nil {
		config.OnCreate(result)
	}

	for i := int32(0); i < result.minWorkers; i++ {
		result.tryAddWorker()
	}

	return result, nil
}

type pool struct {
	queue               chan func()
	queueSize           uint32
	outstanding         int32
	count               int32
	minWorkers          int32
	maxWorkers          int32
	busyWorkers         uint32
	maxIdle             time.Duration
	stopped             atomic.Bool
	externalCloseNotify <-chan struct{}
	closeNotify         chan struct{}
	panicHandler        func(err interface{})
	onWorkCallback      func(workTime time.Duration)
	workF               func(uint32, func())
}

func (self *pool) Queue(work func()) error {
	return self.queueImpl(work, nil)
}

func (self *pool) QueueWithTimeout(work func(), timeout time.Duration) error {
	return self.queueImpl(work, time.After(timeout))
}

func (self *pool) queueImpl(work func(), timeoutC <-chan time.Time) error {
	self.ensureNoStarvation()
	// Count the work as outstanding before it can be picked up by a worker, so a
	// worker can never complete-and-decrement before this increment lands. Undo it
	// on any path where the work isn't actually enqueued.
	self.incrOutstanding()
	select {
	case self.queue <- work:
		self.incrQueueSize()
		self.ensureNoStarvation()
		return nil
	case <-self.closeNotify:
		self.decrOutstanding()
		return errors.Wrap(PoolStoppedError, "cannot queue")
	case <-self.externalCloseNotify:
		self.decrOutstanding()
		return errors.Wrap(PoolStoppedError, "cannot queue, pool stopped externally")
	case <-timeoutC:
		self.decrOutstanding()
		return errors.Wrap(TimeoutError, "cannot queue")
	}
}

func (self *pool) QueueOrError(work func()) error {
	// See queueImpl: count as outstanding before enqueue, undo if not enqueued.
	self.incrOutstanding()
	select {
	case self.queue <- work:
		self.incrQueueSize()
		self.ensureNoStarvation()
		return nil
	case <-self.closeNotify:
		self.decrOutstanding()
		return errors.Wrap(PoolStoppedError, "cannot queue")
	case <-self.externalCloseNotify:
		self.decrOutstanding()
		return errors.Wrap(PoolStoppedError, "cannot queue, pool stopped externally")
	default:
		self.decrOutstanding()
		return errors.Wrap(QueueFullError, "cannot queue")
	}
}

func (self *pool) ensureNoStarvation() {
	if self.minWorkers == 0 && self.GetWorkerCount() == 0 {
		self.tryAddWorker()
	}
}

func (self *pool) Shutdown() {
	if self.stopped.CompareAndSwap(false, true) {
		close(self.closeNotify)
	}
}

func (self *pool) ShutdownAndWait(timeout time.Duration) error {
	self.Shutdown()

	// Wait for every worker to exit, which means all in-flight work has
	// completed. Workers stop pulling from the queue and new workers stop being
	// spawned once the pool is shut down, so the worker count drains to zero.
	return self.awaitCondition(timeout, shutdownPollInterval, "timed out waiting for in-flight work to complete", func() bool {
		return self.GetWorkerCount() == 0
	})
}

func (self *pool) AwaitIdle(timeout time.Duration) error {
	// Wait until no work is outstanding (nothing queued, nothing running).
	return self.awaitCondition(timeout, idlePollInterval, "timed out waiting for pool to become idle", func() bool {
		return self.GetOutstanding() == 0
	})
}

// awaitCondition returns nil once done reports true, or wraps TimeoutError with
// msg if the timeout elapses first. It waits on channels rather than sleeping,
// and the deadline timer fires exactly at the timeout, so the wait doesn't
// overshoot by a poll interval.
func (self *pool) awaitCondition(timeout, pollInterval time.Duration, msg string, done func() bool) error {
	if done() {
		return nil
	}

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-timer.C:
			return errors.Wrap(TimeoutError, msg)
		case <-ticker.C:
			if done() {
				return nil
			}
		}
	}
}

func (self *pool) worker(initialWork func()) {
	defer func() {
		if err := recover(); err != nil {
			if self.panicHandler != nil {
				self.panicHandler(err)
			} else {
				fmt.Printf("panic during pool worker executing (%+v)\n", err)
			}
			self.tryAddWorker()
		}
	}()

	defer func() {
		newCount := self.decrementCount()

		if self.stopped.Load() {
			// Pool is shutting down. Don't respawn workers. The count is still
			// decremented (unlike the running case's respawn logic) so that
			// ShutdownAndWait can observe the pool draining to zero workers.
			return
		}

		// There's a small race condition where the last worker can exit due to idle
		// right as something is queued. If we're the last worker, check again, just
		// to be sure there's nothing queued.
		//
		// There's another race condition where if minWorkers is 1, multiple can exit
		// at the same time and the count can drop to 0. If that happens, start a new
		// worker
		if newCount < self.minWorkers {
			self.addWorkerIfBelowMin()
		} else if newCount == 0 {
			time.AfterFunc(100*time.Millisecond, self.startExtraWorkerIfQueueBusy)
		}
	}()

	if initialWork != nil {
		self.runWork(initialWork)
	}

	// Once shut down, finish the current work (already done by this point)
	// and exit without pulling more from the queue. Queued-but-not-started
	// work is intentionally abandoned.
	for !self.stopped.Load() {
		select {
		case work := <-self.queue:
			self.decrQueueSize()
			self.startExtraWorkerIfQueueBusy()
			self.runWork(work)
		case <-time.After(self.maxIdle):
			if self.getWorkerCount() > self.minWorkers {
				return
			}
		case <-self.closeNotify:
			return
		case <-self.externalCloseNotify:
			self.Shutdown()
			return
		}
	}
}

func (self *pool) startExtraWorkerIfQueueBusy() {
	if self.stopped.Load() {
		return
	}
	if self.getWorkerCount() < self.maxWorkers {
		if workerNumber := self.incrementCount(); workerNumber <= self.maxWorkers {
			select {
			case work := <-self.queue:
				self.decrQueueSize()
				go self.workF(uint32(workerNumber), func() {
					self.worker(work)
				})
			default:
				self.decrementCount()
			}
		} else {
			self.decrementCount()
		}
	}
}

func (self *pool) tryAddWorker() {
	if self.stopped.Load() {
		return
	}
	if self.getWorkerCount() < self.maxWorkers {
		if workerNumber := self.incrementCount(); workerNumber <= self.maxWorkers {
			go self.workF(uint32(workerNumber), func() {
				self.worker(nil)
			})
		} else {
			self.decrementCount()
		}
	}
}

func (self *pool) addWorkerIfBelowMin() {
	if self.stopped.Load() {
		return
	}
	if workerNumber := self.incrementCount(); workerNumber <= self.minWorkers {
		go self.workF(uint32(workerNumber), func() {
			self.worker(nil)
		})
	} else {
		self.decrementCount()
	}
}

func (self *pool) runWork(work func()) {
	self.incrBusyWorkers()
	defer self.decrBusyWorkers()
	defer self.decrOutstanding()
	if self.onWorkCallback != nil {
		start := time.Now()
		work()
		self.onWorkCallback(time.Since(start))
	} else {
		work()
	}
}

func (self *pool) GetWorkerCount() uint32 {
	return uint32(atomic.LoadInt32(&self.count))
}

func (self *pool) getWorkerCount() int32 {
	return atomic.LoadInt32(&self.count)
}

func (self *pool) incrementCount() int32 {
	return atomic.AddInt32(&self.count, 1)
}

func (self *pool) decrementCount() int32 {
	return atomic.AddInt32(&self.count, -1)
}

func (self *pool) GetQueueSize() uint32 {
	return atomic.LoadUint32(&self.queueSize)
}

func (self *pool) incrQueueSize() uint32 {
	return atomic.AddUint32(&self.queueSize, 1)
}

func (self *pool) decrQueueSize() uint32 {
	return atomic.AddUint32(&self.queueSize, ^uint32(0))
}

func (self *pool) GetOutstanding() uint32 {
	// Clamp to zero: the counter is balanced (incremented before enqueue,
	// decremented on completion or enqueue failure), so it should never be
	// negative, but a stray decrement must not surface as a huge uint32.
	if v := atomic.LoadInt32(&self.outstanding); v > 0 {
		return uint32(v)
	}
	return 0
}

func (self *pool) incrOutstanding() int32 {
	return atomic.AddInt32(&self.outstanding, 1)
}

func (self *pool) decrOutstanding() int32 {
	return atomic.AddInt32(&self.outstanding, -1)
}

func (self *pool) GetBusyWorkers() uint32 {
	return atomic.LoadUint32(&self.busyWorkers)
}

func (self *pool) incrBusyWorkers() uint32 {
	return atomic.AddUint32(&self.busyWorkers, 1)
}

func (self *pool) decrBusyWorkers() uint32 {
	return atomic.AddUint32(&self.busyWorkers, ^uint32(0))
}
