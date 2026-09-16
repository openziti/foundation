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

import "sync/atomic"

// AtomicBitSet is a set of up to 32 flags updated atomically as one word. All methods are safe
// for concurrent use. The zero value is an empty set.
type AtomicBitSet uint32

// BitSet is a snapshot of an AtomicBitSet's word, as returned by GetAndSet and GetAndClear. It is
// a plain value and is not safe to share between goroutines that mutate it.
type BitSet uint32

// IsSet reports whether the bit at index is set in the snapshot.
func (self BitSet) IsSet(index int) bool {
	return isBitSetAtIndex(uint32(self), index)
}

// Set sets or clears the bit at index.
func (self *AtomicBitSet) Set(index int, val bool) {
	if val {
		self.GetAndSet(index)
	} else {
		self.GetAndClear(index)
	}
}

// GetAndSet sets the bit at index and returns the set as it was immediately before, in a single
// atomic read-modify-write. Two goroutines each setting their own bit this way are guaranteed
// that at least one of them sees the other's bit in the returned value, which a separate store
// and load cannot promise without ordering both sides.
func (self *AtomicBitSet) GetAndSet(index int) BitSet {
	return BitSet(atomic.OrUint32((*uint32)(self), 1<<index))
}

// GetAndClear clears the bit at index and returns the set as it was immediately before, in a
// single atomic read-modify-write.
func (self *AtomicBitSet) GetAndClear(index int) BitSet {
	return BitSet(atomic.AndUint32((*uint32)(self), ^uint32(1<<index)))
}

func (self *AtomicBitSet) IsSet(index int) bool {
	return isBitSetAtIndex(self.Load(), index)
}

func (self *AtomicBitSet) CompareAndSet(index int, current, next bool) bool {
	for {
		currentSet := self.Load()
		if isBitSetAtIndex(currentSet, index) != current {
			return false
		}
		nextSet := setBitAtIndex(currentSet, index, next)
		if self.CompareAndSetAll(currentSet, nextSet) {
			return true
		}
	}
}

func (self *AtomicBitSet) Store(val uint32) {
	atomic.StoreUint32((*uint32)(self), val)
}

func (self *AtomicBitSet) Load() uint32 {
	return atomic.LoadUint32((*uint32)(self))
}

func (self *AtomicBitSet) CompareAndSetAll(current, next uint32) bool {
	return atomic.CompareAndSwapUint32((*uint32)(self), current, next)
}

func setBitAtIndex(bitset uint32, index int, val bool) uint32 {
	if val {
		return bitset | (1 << index)
	}
	return bitset & ^(1 << index)
}

func isBitSetAtIndex(bitset uint32, index int) bool {
	return bitset&(1<<index) != 0
}

func (self *AtomicBitSet) SetBits(startIndex, width int, val uint32) {
	mask := uint32(((1 << width) - 1) << startIndex)
	for {
		current := self.Load()
		next := (current &^ mask) | ((val << startIndex) & mask)
		if self.CompareAndSetAll(current, next) {
			return
		}
	}
}

func (self *AtomicBitSet) GetBits(startIndex, width int) uint32 {
	mask := uint32(((1 << width) - 1) << startIndex)
	return (self.Load() & mask) >> startIndex
}
