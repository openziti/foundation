package concurrenz

import "time"

type Semaphore interface {
	Acquire()
	AcquireWithTimeout(t time.Duration) bool
	TryAcquire() bool
	Release() bool
}

func NewSemaphore(size int) Semaphore {
	result := &semaphoreImpl{
		c: make(chan struct{}, size),
	}
	for result.Release() {
	}
	return result
}

type semaphoreImpl struct {
	c chan struct{}
}

func (self *semaphoreImpl) Acquire() {
	<-self.c
}

func (self *semaphoreImpl) AcquireWithTimeout(t time.Duration) bool {
	select {
	case <-self.c:
		return true
	case <-time.After(t):
		return false
	}
}

func (self *semaphoreImpl) TryAcquire() bool {
	select {
	case <-self.c:
		return true
	default:
		return false
	}
}

func (self *semaphoreImpl) Release() bool {
	select {
	case self.c <- struct{}{}:
		return true
	default:
		return false
	}
}
