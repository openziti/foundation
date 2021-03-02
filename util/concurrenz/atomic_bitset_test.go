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
