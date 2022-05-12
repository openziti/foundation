package goroutines

import (
	"fmt"
	"github.com/openziti/foundation/util/concurrenz"
	"github.com/pkg/errors"
	"sync/atomic"
	"time"
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

// Pool represents a goroutine worker pool that can be configured with a queue size and min and max sizes.
//      The pool will start with min size goroutines and will add more if the queue isn't staying empty.
//      After a worker has been idle for a configured time, it will stop
type Pool interface {
	// Queue submits a unit of work to the pool. It will return an error if the pool is shutdown
	Queue(func()) error

	// QueueWithTimeout submits a unit of work to the pool. It will return an error if the pool is shutdown or
	// if the work cannot be submitted to the work queue before the given timeout elapses
	QueueWithTimeout(func(), time.Duration) error

	// QueueOrError submits a unit of work to the pool. It will return an error if the pool is shutdown or
	// if the work cannot be submitted to the work queue immediately
	QueueOrError(func()) error

	// Shutdown stops all workers as they finish work and prevents new work from being submitted to the queue
	Shutdown()
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
	result := &pool{
		queue:               make(chan func(), int(config.QueueSize)),
		minWorkers:          config.MinWorkers,
		maxWorkers:          config.MaxWorkers,
		maxIdle:             config.IdleTime,
		externalCloseNotify: config.CloseNotify,
		closeNotify:         make(chan struct{}),
		panicHandler:        config.PanicHandler,
	}
	for result.count < result.minWorkers {
		result.tryAddWorker()
	}
	return result, nil
}

type pool struct {
	queue               chan func()
	count               uint32
	minWorkers          uint32
	maxWorkers          uint32
	maxIdle             time.Duration
	stopped             concurrenz.AtomicBoolean
	externalCloseNotify <-chan struct{}
	closeNotify         chan struct{}
	panicHandler        func(err interface{})
}

func (self *pool) Queue(work func()) error {
	return self.queueImpl(work, nil)
}

func (self *pool) QueueWithTimeout(work func(), timeout time.Duration) error {
	return self.queueImpl(work, time.After(timeout))
}

func (self *pool) queueImpl(work func(), timeoutC <-chan time.Time) error {
	if self.getCount() == 0 {
		self.tryAddWorker()
	}

	select {
	case self.queue <- work:
		return nil
	case <-self.closeNotify:
		return errors.Wrap(PoolStoppedError, "cannot queue")
	case <-self.externalCloseNotify:
		return errors.Wrap(PoolStoppedError, "cannot queue, pool stopped externally")
	case <-timeoutC:
		return errors.Wrap(TimeoutError, "cannot queue")
	}
}

func (self *pool) QueueOrError(work func()) error {
	if self.getCount() == 0 {
		self.tryAddWorker()
	}

	select {
	case self.queue <- work:
		return nil
	case <-self.closeNotify:
		return errors.Wrap(PoolStoppedError, "cannot queue")
	case <-self.externalCloseNotify:
		return errors.Wrap(PoolStoppedError, "cannot queue, pool stopped externally")
	default:
		return errors.Wrap(QueueFullError, "cannot queue")
	}
}

func (self *pool) Shutdown() {
	if self.stopped.CompareAndSwap(false, true) {
		close(self.closeNotify)
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
		self.decrementCount()
	}()

	if initialWork != nil {
		initialWork()
		initialWork = nil
	}

	for {
		select {
		case work := <-self.queue:
			self.startExtraWorkerIfQueueBusy()
			work()
		case <-time.After(self.maxIdle):
			if self.getCount() > self.minWorkers {
				return
			}
		case <-self.closeNotify:
			return
		case <-self.externalCloseNotify:
			return
		}
	}
}

func (self *pool) startExtraWorkerIfQueueBusy() {
	if self.getCount() < self.maxWorkers {
		if self.incrementCount() <= self.maxWorkers {
			select {
			case work := <-self.queue:
				go self.worker(work)
			default:
				self.decrementCount()
			}
		} else {
			self.decrementCount()
		}
	}
}

func (self *pool) tryAddWorker() {
	if self.getCount() < self.maxWorkers {
		if self.incrementCount() <= self.maxWorkers {
			go self.worker(nil)
		} else {
			self.decrementCount()
		}
	}
}

func (self *pool) getCount() uint32 {
	return atomic.LoadUint32(&self.count)
}

func (self *pool) incrementCount() uint32 {
	return atomic.AddUint32(&self.count, 1)
}

func (self *pool) decrementCount() uint32 {
	return atomic.AddUint32(&self.count, ^uint32(0))
}
