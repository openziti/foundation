/*
	Copyright NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package concurrenz

import "sync"

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: map[K]V{},
	}
}

type SyncMap[K comparable, V any] struct {
	l sync.RWMutex
	m map[K]V
}

func (self *SyncMap[K, V]) Put(k K, v V) {
	self.l.Lock()
	self.m[k] = v
	self.l.Unlock()
}

func (self *SyncMap[K, V]) Remove(k K) {
	self.l.Lock()
	delete(self.m, k)
	self.l.Unlock()
}

func (self *SyncMap[K, V]) Len() int {
	self.l.RLock()
	defer self.l.RUnlock()
	return len(self.m)
}

func (self *SyncMap[K, V]) Get(k K) (V, bool) {
	self.l.RLock()
	defer self.l.RUnlock()
	v, ok := self.m[k]
	return v, ok
}

func (self *SyncMap[K, V]) RangeAll(f func(K, V)) {
	self.l.RLock()
	defer self.l.RUnlock()
	for k, v := range self.m {
		f(k, v)
	}
}

func (self *SyncMap[K, V]) Range(f func(K, V) bool) {
	self.l.RLock()
	defer self.l.RUnlock()
	for k, v := range self.m {
		if !f(k, v) {
			break
		}
	}
}

func (self *SyncMap[K, V]) Clear() {
	self.l.Lock()
	self.m = make(map[K]V)
	self.l.Unlock()
}

func (self *SyncMap[K, V]) WithReadLock(f func(map[K]V)) {
	self.l.RLock()
	f(self.m)
	self.l.RUnlock()
}

func (self *SyncMap[K, V]) WithWriteLock(f func(map[K]V)) {
	self.l.Lock()
	f(self.m)
	self.l.Unlock()
}

func (self *SyncMap[K, V]) ToMap() map[K]V {
	result := map[K]V{}
	self.RangeAll(func(k K, v V) {
		result[k] = v
	})
	return result
}
