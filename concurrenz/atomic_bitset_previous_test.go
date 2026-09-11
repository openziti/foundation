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

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AtomicBitset_SetAndClearAndGetPrevious(t *testing.T) {
	req := require.New(t)
	var bits AtomicBitSet

	prev := bits.SetAndGetPrevious(3)
	req.False(prev.IsSet(3))
	req.True(bits.IsSet(3))

	prev = bits.SetAndGetPrevious(3)
	req.True(prev.IsSet(3), "setting an already set bit reports it as previously set")
	req.True(bits.IsSet(3))

	prev = bits.SetAndGetPrevious(7)
	req.True(prev.IsSet(3), "the previous value carries unrelated bits")
	req.False(prev.IsSet(7))
	req.True(bits.IsSet(3))
	req.True(bits.IsSet(7))

	prev = bits.ClearAndGetPrevious(3)
	req.True(prev.IsSet(3))
	req.True(prev.IsSet(7))
	req.False(bits.IsSet(3))
	req.True(bits.IsSet(7), "clearing one bit leaves the others")

	prev = bits.ClearAndGetPrevious(3)
	req.False(prev.IsSet(3), "clearing an already clear bit reports it as previously clear")

	prev = bits.SetAndGetPrevious(31)
	req.False(prev.IsSet(31))
	req.True(bits.IsSet(31))
	req.Equal(uint32(1<<31|1<<7), bits.Load())
}

// Two goroutines each set their own bit and look for the other's in the previous value. Because
// each operation is a single read-modify-write, at least one of them must see the other's bit;
// a separate store and load offers no such guarantee. Iterated to give the race a chance.
func Test_AtomicBitset_SetAndGetPrevious_OneSideAlwaysSeesTheOther(t *testing.T) {
	req := require.New(t)

	const first, second = 0, 1
	for i := 0; i < 20000; i++ {
		var bits AtomicBitSet
		var start, done sync.WaitGroup
		start.Add(1)
		done.Add(2)

		var firstSawSecond, secondSawFirst bool
		go func() {
			defer done.Done()
			start.Wait()
			firstSawSecond = bits.SetAndGetPrevious(first).IsSet(second)
		}()
		go func() {
			defer done.Done()
			start.Wait()
			secondSawFirst = bits.SetAndGetPrevious(second).IsSet(first)
		}()
		start.Done()
		done.Wait()

		req.True(firstSawSecond || secondSawFirst, "iteration %d: neither side saw the other", i)
		req.False(firstSawSecond && secondSawFirst, "iteration %d: both sides cannot have gone second", i)
	}
}
