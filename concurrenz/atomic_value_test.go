package concurrenz

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAtomicValue_CompareAndSwap(t *testing.T) {
	v := &AtomicValue[string]{}

	req := require.New(t)
	req.Equal("", v.Load())
	v.Store("hello")
	req.Equal("hello", v.Load())
	req.Equal(false, v.CompareAndSwap("world", "foo"))
	req.Equal("hello", v.Load())
	req.Equal(true, v.CompareAndSwap("hello", "world"))
	req.Equal("world", v.Load())
	v.Store("foo")
	req.Equal("foo", v.Load())
}
