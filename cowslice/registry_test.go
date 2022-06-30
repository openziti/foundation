package cowslice

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testHandler interface {
	Accept(val struct{})
}

type testHandlerImpl struct{}

func (t testHandlerImpl) Accept(val struct{}) {
	panic("implement me")
}

type testHandlerImpl2 struct{}

func (t testHandlerImpl2) Accept(val struct{}) {
	panic("implement me")
}

func TestCowSlice_Append(t *testing.T) {
	req := require.New(t)
	cw := NewCowSlice(make([]testHandler, 0))
	req.NotNil(cw.Value())

	h := &testHandlerImpl{}
	Append(cw, h)

	cur := cw.Value()
	req.NotNil(cur)
	tc := cur.([]testHandler)
	req.Equal(1, len(tc))
	req.Equal(h, tc[0])

	h2 := &testHandlerImpl2{}
	Append(cw, h2)

	cur = cw.Value()
	req.NotNil(cur)
	tc = cur.([]testHandler)
	req.Equal(2, len(tc))
	req.Equal(h, tc[0])
	req.Equal(h2, tc[1])

	Delete(cw, h)
	cur = cw.Value()
	req.NotNil(cur)
	tc = cur.([]testHandler)
	req.Equal(1, len(tc))
	req.Equal(h2, tc[0])

	Delete(cw, h2)
	cur = cw.Value()
	req.NotNil(cur)
	tc = cur.([]testHandler)
	req.Equal(0, len(tc))
}

func TestCloseSlice_Delete(t *testing.T) {

	t.Run("adding and removing a single element succeeds", func(t *testing.T) {
		req := require.New(t)
		cw := NewCowSlice(make([]testHandler, 0))

		h := &testHandlerImpl{}
		Append(cw, h)
		Delete(cw, h)

		cur := cw.Value()
		req.NotNil(cur)

		tc, ok := cur.([]testHandler)
		req.True(ok)
		req.Equal(0, len(tc))
	})

	t.Run("delete on empty cow slice does not panic", func(t *testing.T) {
		req := require.New(t)
		cw := NewCowSlice(make([]testHandler, 0))

		h := &testHandlerImpl{}
		Delete(cw, h)

		cur := cw.Value()
		req.NotNil(cur)

		tc, ok := cur.([]testHandler)
		req.True(ok)
		req.Equal(0, len(tc))
	})
}
