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
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_AtomicBitset(t *testing.T) {
	req := require.New(t)

	bs := AtomicBitSet(0)
	bs.Set(0, true)
	bs.Set(2, true)
	bs.Set(4, true)

	req.True(bs.IsSet(0))
	req.False(bs.IsSet(1))
	req.True(bs.IsSet(2))
	req.False(bs.IsSet(3))
	req.True(bs.IsSet(4))
	req.False(bs.IsSet(5))

	bs.Set(0, false)
	bs.Set(2, false)
	bs.Set(4, false)
	bs.Set(1, true)
	bs.Set(3, true)
	bs.Set(5, true)

	req.False(bs.IsSet(0))
	req.True(bs.IsSet(1))
	req.False(bs.IsSet(2))
	req.True(bs.IsSet(3))
	req.False(bs.IsSet(4))
	req.True(bs.IsSet(5))

	req.False(bs.CompareAndSet(0, true, false))
	req.False(bs.IsSet(0))
	req.True(bs.CompareAndSet(0, false, true))
	req.True(bs.IsSet(0))
}

func Test_AtomicBitset_MultiBit(t *testing.T) {
	req := require.New(t)

	bs := AtomicBitSet(0)

	// Store a 2-bit value at position 6
	bs.SetBits(6, 2, 3)
	req.Equal(uint32(3), bs.GetBits(6, 2))

	// Overwrite with a different value
	bs.SetBits(6, 2, 1)
	req.Equal(uint32(1), bs.GetBits(6, 2))

	// Zero it out
	bs.SetBits(6, 2, 0)
	req.Equal(uint32(0), bs.GetBits(6, 2))

	// Ensure other bits aren't affected
	bs.Set(0, true)
	bs.Set(5, true)
	bs.SetBits(6, 2, 2)
	req.True(bs.IsSet(0))
	req.True(bs.IsSet(5))
	req.Equal(uint32(2), bs.GetBits(6, 2))

	// Test wider field (3 bits at position 2)
	bs2 := AtomicBitSet(0)
	bs2.SetBits(2, 3, 5) // binary 101
	req.Equal(uint32(5), bs2.GetBits(2, 3))
	req.False(bs2.IsSet(0))
	req.False(bs2.IsSet(1))
	req.True(bs2.IsSet(2))  // bit 2 = 1
	req.False(bs2.IsSet(3)) // bit 3 = 0
	req.True(bs2.IsSet(4))  // bit 4 = 1
}
