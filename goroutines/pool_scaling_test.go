package goroutines

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPoolGrowsWhileWorkersBusy(t *testing.T) {
	req := require.New(t)

	p, err := NewPool(PoolConfig{
		QueueSize:  20,
		MinWorkers: 0,
		MaxWorkers: 8,
		IdleTime:   time.Second,
	})
	req.NoError(err)
	defer p.Shutdown()

	release := make(chan struct{})
	var done sync.WaitGroup
	for i := 0; i < 20; i++ {
		done.Add(1)
		req.NoError(p.QueueOrError(func() {
			defer done.Done()
			<-release
		}))
	}

	req.EventuallyWithT(func(c *assert.CollectT) {
		assert.Equal(c, uint32(8), p.GetBusyWorkers(), "every worker up to max should be running while the queue holds work")
	}, time.Second, time.Millisecond)
	req.Equal(uint32(8), p.GetWorkerCount())
	req.Equal(12, int(p.GetQueueSize()))

	close(release)
	done.Wait()
}

func TestPoolWithQueueSizeOneGrows(t *testing.T) {
	req := require.New(t)

	p, err := NewPool(PoolConfig{
		QueueSize:  1,
		MinWorkers: 0,
		MaxWorkers: 4,
		IdleTime:   time.Second,
	})
	req.NoError(err)
	defer p.Shutdown()

	release := make(chan struct{})
	var done sync.WaitGroup
	queueBlocked := func() error {
		done.Add(1)
		err := p.QueueOrError(func() {
			defer done.Done()
			<-release
		})
		if err != nil {
			done.Done()
		}
		return err
	}

	for i := 1; i <= 4; i++ {
		req.NoError(queueBlocked())
		req.EventuallyWithT(func(c *assert.CollectT) {
			assert.Equal(c, uint32(i), p.GetBusyWorkers(), "work %d should start on a new worker", i)
		}, time.Second, time.Millisecond)
	}

	// At max workers the queue fills, and then QueueOrError reports it full.
	req.NoError(queueBlocked())
	req.ErrorIs(queueBlocked(), QueueFullError)
	req.Equal(uint32(4), p.GetWorkerCount())

	close(release)
	done.Wait()
}

// TestPoolDoesNotStrandWorkAsWorkersExit has one or more producers each queue work one item at a time, with
// gaps around the idle time, so workers keep exiting, several at once with several producers, as new work
// arrives. Every item must still run.
func TestPoolDoesNotStrandWorkAsWorkersExit(t *testing.T) {
	for _, minWorkers := range []uint32{0, 1} {
		for _, producers := range []int{1, 4} {
			t.Run(fmt.Sprintf("min workers %d, producers %d", minWorkers, producers), func(t *testing.T) {
				req := require.New(t)

				p, err := NewPool(PoolConfig{
					QueueSize:  uint32(producers),
					MinWorkers: minWorkers,
					MaxWorkers: 8,
					IdleTime:   time.Millisecond,
				})
				req.NoError(err)
				defer p.Shutdown()

				errs := make(chan error, producers)
				for i := 0; i < producers; i++ {
					go func() { errs <- queueEachAndWait(p, 2000) }()
				}
				for i := 0; i < producers; i++ {
					req.NoError(<-errs)
				}

				req.EventuallyWithT(func(c *assert.CollectT) {
					assert.Equal(c, minWorkers, p.GetWorkerCount(), "workers should settle at min")
				}, time.Second, time.Millisecond)
			})
		}
	}
}

// queueEachAndWait queues count items one at a time, alternating Queue and QueueOrError, and waits for each
// to run before sleeping up to 2ms. It returns an error if an item can't be queued or doesn't run within a
// second.
func queueEachAndWait(p Pool, count int) error {
	for i := 0; i < count; i++ {
		ran := make(chan struct{})
		work := func() { close(ran) }

		var err error
		if i%2 == 0 {
			err = p.Queue(work)
		} else {
			err = p.QueueOrError(work)
		}
		if err != nil {
			return fmt.Errorf("item %d: %w", i, err)
		}

		select {
		case <-ran:
		case <-time.After(time.Second):
			return fmt.Errorf("work stranded: item %d never ran, workers: %d, queued: %d",
				i, p.GetWorkerCount(), p.GetQueueSize())
		}

		time.Sleep(time.Duration(rand.Int63n(int64(2 * time.Millisecond))))
	}
	return nil
}

// TestPoolKeepsMinWorkersWhenWorkersExitTogether repeatedly grows the pool to max, then lets every worker go
// idle at the same moment, so their idle timeouts fire together. Each round must run all of its work at once,
// and the pool must settle back at min workers, not below it.
func TestPoolKeepsMinWorkersWhenWorkersExitTogether(t *testing.T) {
	req := require.New(t)

	p, err := NewPool(PoolConfig{
		QueueSize:  16,
		MinWorkers: 1,
		MaxWorkers: 16,
		IdleTime:   5 * time.Millisecond,
	})
	req.NoError(err)
	defer p.Shutdown()

	for round := 0; round < 200; round++ {
		release := make(chan struct{})
		var running atomic.Int32
		var done sync.WaitGroup
		for i := 0; i < 16; i++ {
			done.Add(1)
			req.NoError(p.QueueOrError(func() {
				defer done.Done()
				running.Add(1)
				<-release
			}))
		}
		req.EventuallyWithT(func(c *assert.CollectT) {
			assert.Equal(c, int32(16), running.Load(), "round %d: all work should run at once, workers: %d, queued: %d",
				round, p.GetWorkerCount(), p.GetQueueSize())
		}, time.Second, time.Millisecond)
		close(release)
		done.Wait()

		req.EventuallyWithT(func(c *assert.CollectT) {
			assert.Equal(c, uint32(1), p.GetWorkerCount(), "round %d: workers should settle at min", round)
		}, time.Second, time.Millisecond)
		time.Sleep(10 * time.Millisecond)
		req.Equal(uint32(1), p.GetWorkerCount(), "round %d: workers dropped below min", round)
	}
}

// TestPoolWorkersExitingTogether holds every worker that has decided to exit until all of them have, so
// they leave together. The pool must keep min workers, and must run work queued while they were leaving.
func TestPoolWorkersExitingTogether(t *testing.T) {
	const workers = 4
	for _, minWorkers := range []uint32{0, 1} {
		for _, queueWhileExiting := range []bool{false, true} {
			t.Run(fmt.Sprintf("min workers %d, queue while exiting %v", minWorkers, queueWhileExiting), func(t *testing.T) {
				req := require.New(t)

				var hold atomic.Bool
				arrived := make(chan struct{}, workers)
				proceed := make(chan struct{})
				p, err := NewPool(PoolConfig{
					QueueSize:  workers,
					MinWorkers: minWorkers,
					MaxWorkers: workers,
					IdleTime:   10 * time.Millisecond,
					onIdleExit: func() {
						if hold.Load() {
							arrived <- struct{}{}
							<-proceed
						}
					},
				})
				req.NoError(err)
				defer p.Shutdown()

				release := make(chan struct{})
				var running atomic.Int32
				for i := 0; i < workers; i++ {
					req.NoError(p.QueueOrError(func() {
						running.Add(1)
						<-release
					}))
				}
				req.EventuallyWithT(func(c *assert.CollectT) {
					assert.Equal(c, int32(workers), running.Load(), "the pool should grow to max")
				}, time.Second, time.Millisecond)
				hold.Store(true)
				close(release)

				// Held workers stay in the count, so each one sees it above min and decides to exit.
				for i := 0; i < workers; i++ {
					select {
					case <-arrived:
					case <-time.After(time.Second):
						req.FailNow("worker didn't decide to exit", "workers arrived: %d of %d", i, workers)
					}
				}

				ran := make(chan struct{})
				if queueWhileExiting {
					req.NoError(p.QueueOrError(func() { close(ran) }))
				}
				hold.Store(false)
				close(proceed)

				if queueWhileExiting {
					select {
					case <-ran:
					case <-time.After(time.Second):
						req.FailNow("work stranded", "workers: %d, queued: %d", p.GetWorkerCount(), p.GetQueueSize())
					}
				}

				req.EventuallyWithT(func(c *assert.CollectT) {
					assert.Equal(c, minWorkers, p.GetWorkerCount(), "workers should settle at min")
				}, time.Second, time.Millisecond)
				time.Sleep(30 * time.Millisecond)
				req.Equal(minWorkers, p.GetWorkerCount(), "workers dropped below min")
			})
		}
	}
}

func TestPoolDoesNotOvershootBacklog(t *testing.T) {
	req := require.New(t)

	p, err := NewPool(PoolConfig{
		QueueSize:  10,
		MinWorkers: 0,
		MaxWorkers: 10,
		IdleTime:   time.Second,
	})
	req.NoError(err)
	defer p.Shutdown()

	release := make(chan struct{})
	var done sync.WaitGroup
	for i := 0; i < 3; i++ {
		done.Add(1)
		req.NoError(p.QueueOrError(func() {
			defer done.Done()
			<-release
		}))
	}

	req.EventuallyWithT(func(c *assert.CollectT) {
		assert.Equal(c, uint32(3), p.GetBusyWorkers())
	}, time.Second, time.Millisecond)
	// Give any surplus worker a chance to appear before checking there isn't one.
	time.Sleep(20 * time.Millisecond)
	req.Equal(uint32(3), p.GetWorkerCount())

	close(release)
	done.Wait()
}
