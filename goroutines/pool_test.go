package goroutines

import (
	"errors"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	concurrenz2 "github.com/openziti/foundation/v2/concurrenz"
	"github.com/stretchr/testify/require"
)

const maxIterations = 10

func TestPoolWithMinTwo(t *testing.T) {
	val, err := NewPool(PoolConfig{
		QueueSize:   100,
		MinWorkers:  2,
		MaxWorkers:  10,
		IdleTime:    100 * time.Millisecond,
		CloseNotify: nil,
		PanicHandler: func(err interface{}) {
			fmt.Printf("panic: %v\n", err)
		},
	})

	require.NoError(t, err)
	p := val.(*pool)

	for i := 0; i < maxIterations; i++ {
		RunPoolTest(t, p)
	}
}

func TestPoolWithMinZero(t *testing.T) {
	val, err := NewPool(PoolConfig{
		QueueSize:   100,
		MinWorkers:  0,
		MaxWorkers:  10,
		IdleTime:    100 * time.Millisecond,
		CloseNotify: nil,
		PanicHandler: func(err interface{}) {
			fmt.Printf("panic: %v\n", err)
		},
	})

	require.NoError(t, err)
	p := val.(*pool)

	for i := 0; i < maxIterations; i++ {
		RunPoolTest(t, p)
	}
}

func TestPoolWithMinOne(t *testing.T) {
	val, err := NewPool(PoolConfig{
		QueueSize:   100,
		MinWorkers:  1,
		MaxWorkers:  10,
		IdleTime:    100 * time.Millisecond,
		CloseNotify: nil,
		PanicHandler: func(err interface{}) {
			fmt.Printf("panic: %v\n", err)
		},
	})

	require.NoError(t, err)
	p := val.(*pool)

	for i := 0; i < maxIterations; i++ {
		RunPoolTest(t, p)
	}
}

func RunPoolTest(t *testing.T, p *pool) {
	req := require.New(t)
	busyWork := &poolBusier{workPool: p}

	req.Equal(int(p.minWorkers), int(p.GetWorkerCount()))

	t.Run("test 2 workers", func(t *testing.T) {
		busyWork.KeepBusy(2, 0)
		time.Sleep(50 * time.Millisecond)
		count := p.GetWorkerCount()
		req.True(count == 2 || count == 3, "count should be within 1 of min. was %v", count)
		req.NoError(busyWork.CloseAndWait())
		time.Sleep(5 * time.Millisecond)

		time.Sleep(150 * time.Millisecond)
		req.Equal(int(p.minWorkers), int(p.GetWorkerCount()))
	})

	t.Run("test 8 workers", func(t *testing.T) {
		busyWork.KeepBusy(8, 0)

		for i := 0; i < 10; i++ {
			if p.GetWorkerCount() == 8 {
				break
			}
			time.Sleep(10 * time.Millisecond)
		}

		count := p.GetWorkerCount()
		req.True(count >= 7 && count <= 9, "count should be within 1 of 8 was %v", count)
		req.NoError(busyWork.CloseAndWait())

		time.Sleep(50 * time.Millisecond)
		req.True(count >= 7 && count <= 9, "count should be within 1 of 8 was %v", count)

		time.Sleep(150 * time.Millisecond)
		req.Equal(int(p.minWorkers), int(p.GetWorkerCount()))
	})

	t.Run("test busy queue", func(t *testing.T) {
		busyWork.KeepBusy(15, 0)
		for i := 0; i < 10; i++ {
			if p.GetWorkerCount() == 10 {
				break
			}
			time.Sleep(10 * time.Millisecond)
		}
		req.Equal(10, int(p.GetWorkerCount()))
		req.NoError(busyWork.CloseAndWait())

		time.Sleep(150 * time.Millisecond)
		req.Equal(int(p.minWorkers), int(p.GetWorkerCount()))
	})

	t.Run("test busy queue with panics", func(t *testing.T) {
		busyWork.KeepBusy(15, 12)
		time.Sleep(50 * time.Millisecond)
		req.Equal(10, int(p.GetWorkerCount()))
		req.NoError(busyWork.CloseAndWait())

		time.Sleep(150 * time.Millisecond)
		req.Equal(int(p.minWorkers), int(p.GetWorkerCount()))
	})
}

func TestQueueOrError(t *testing.T) {
	val, err := NewPool(PoolConfig{
		QueueSize:   1,
		MinWorkers:  1,
		MaxWorkers:  1,
		IdleTime:    100 * time.Millisecond,
		CloseNotify: nil,
		PanicHandler: func(err interface{}) {
			fmt.Printf("panic: %v\n", err)
		},
	})
	req := require.New(t)
	req.NoError(err)

	running := make(chan struct{})

	err = val.QueueOrError(func() {
		close(running)
		time.Sleep(100 * time.Millisecond)
	})
	req.NoError(err)

	select {
	case <-running:
	case <-time.After(time.Second):
		req.FailNow("timed out waiting for first task to run")
	}

	err = val.QueueOrError(func() {
		time.Sleep(100 * time.Millisecond)
	})
	req.NoError(err)

	err = val.QueueOrError(func() {
		time.Sleep(100 * time.Millisecond)
	})
	req.Error(err)
	req.ErrorIs(err, QueueFullError)
}

type poolBusier struct {
	workPool Pool
	stopped  atomic.Bool
	errorC   chan error
	done     chan struct{}
}

func (self *poolBusier) KeepBusy(count int, panicCount int) {
	self.stopped.Store(false)
	self.done = make(chan struct{})
	self.errorC = make(chan error, count)
	go func() {
		defer close(self.done)
		sema := concurrenz2.NewSemaphore(count)
		for !self.stopped.Load() {
			sema.Acquire()
			var err error
			if panicCount > 0 {
				panicCount--
				err = self.workPool.Queue(func() {
					sema.Release()
					panic(errors.New("panic"))
				})
			} else {
				err = self.workPool.Queue(func() {
					time.Sleep(20 * time.Millisecond)
					sema.Release()
				})
			}
			if err != nil {
				self.errorC <- err
				return
			}
		}
	}()
}

func (self *poolBusier) CloseAndWait() error {
	self.stopped.Store(true)
	<-self.done
	select {
	case err := <-self.errorC:
		return err
	default:
		return nil
	}
}
