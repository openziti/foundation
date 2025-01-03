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

func NewSyncSet[K comparable]() *SyncSet[K] {
	return &SyncSet[K]{
		m: map[K]struct{}{},
	}
}

type SyncSet[K comparable] struct {
	l sync.RWMutex
	m map[K]struct{}
}

func (self *SyncSet[K]) Add(k K) {
	self.l.Lock()
	self.m[k] = struct{}{}
	self.l.Unlock()
}

func (self *SyncSet[K]) Remove(k K) {
	self.l.Lock()
	delete(self.m, k)
	self.l.Unlock()
}

func (self *SyncSet[K]) Len() int {
	self.l.RLock()
	defer self.l.RUnlock()
	return len(self.m)
}

func (self *SyncSet[K]) Contains(k K) bool {
	self.l.RLock()
	defer self.l.RUnlock()
	_, ok := self.m[k]
	return ok
}

func (self *SyncSet[K]) RangeAll(f func(K)) {
	self.l.RLock()
	defer self.l.RUnlock()
	for k := range self.m {
		f(k)
	}
}

func (self *SyncSet[K]) Range(f func(K) bool) {
	self.l.RLock()
	defer self.l.RUnlock()
	for k := range self.m {
		if !f(k) {
			break
		}
	}
}

func (self *SyncSet[K]) Clear() {
	self.l.Lock()
	self.m = make(map[K]struct{})
	self.l.Unlock()
}

func (self *SyncSet[K]) WithWriteLock(f func(map[K]struct{})) {
	self.l.Lock()
	f(self.m)
	self.l.Unlock()
}

func (self *SyncSet[K]) WithReadLock(f func(map[K]struct{})) {
	self.l.RLock()
	f(self.m)
	self.l.RUnlock()
}

func (self *SyncSet[K]) ToMap() map[K]struct{} {
	result := map[K]struct{}{}
	self.RangeAll(func(k K) {
		result[k] = struct{}{}
	})
	return result
}
