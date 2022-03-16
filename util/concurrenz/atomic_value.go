package concurrenz

import "sync/atomic"

type AtomicValue[T any] atomic.Value

func (self *AtomicValue[T]) Store(val T) {
	(*atomic.Value)(self).Store(val)
}

func (self *AtomicValue[T]) Load() T {
	var result T
	if val := (*atomic.Value)(self).Load(); val != nil {
		result = val.(T)
	}
	return result
}
