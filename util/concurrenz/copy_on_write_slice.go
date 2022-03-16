package concurrenz

import (
	"sync"
)

type CopyOnWriteSlice[T any] struct {
	value AtomicValue[[]T]
	lock  sync.Mutex
}

func (self *CopyOnWriteSlice[T]) Value() []T {
	return self.value.Load()
}

func (self *CopyOnWriteSlice[T]) Append(toAdd T) {
	self.lock.Lock()
	defer self.lock.Unlock()

	currentSlice := self.value.Load()
	newSlice := append(currentSlice, toAdd)
	self.value.Store(newSlice)
}

func (self *CopyOnWriteSlice[T]) Delete(toRemove T) {
	self.lock.Lock()
	defer self.lock.Unlock()

	currentSlice := self.value.Load()
	newSlice := make([]T, 0, len(currentSlice))
	for _, val := range currentSlice {
		if val != toRemove {
			newSlice = append(newSlice, val)
		}
	}
	self.value.Store(newSlice)
}
