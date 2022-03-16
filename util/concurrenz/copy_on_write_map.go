package concurrenz

import "sync"

type CopyOnWriteMap[K comparable, V any] struct {
	value AtomicValue[map[K]V]
	lock  sync.Mutex
}

func (self *CopyOnWriteMap[K, V]) Put(key K, value V) {
	self.lock.Lock()
	defer self.lock.Unlock()

	var current = self.value.Load()
	mapCopy := map[K]V{}
	for k, v := range current {
		mapCopy[k] = v
	}
	mapCopy[key] = value
	self.value.Store(mapCopy)
}

func (self *CopyOnWriteMap[K, V]) Get(key K) V {
	return self.value.Load()[key]
}

func (self *CopyOnWriteMap[K, V]) AsMap() map[K]V {
	return self.value.Load()
}
