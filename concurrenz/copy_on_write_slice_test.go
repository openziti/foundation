package concurrenz

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testHandler interface {
	Accept(val struct{})
}

type testHandlerImpl struct{}

func (t testHandlerImpl) Accept(struct{}) {
	panic("implement me")
}

type testHandlerImpl2 struct{}

func (t testHandlerImpl2) Accept(struct{}) {
	panic("implement me")
}

func TestCowSlice_Append(t *testing.T) {
	req := require.New(t)
	cw := &CopyOnWriteSlice[testHandler]{}
	req.Equal(0, len(cw.Value()))

	h := &testHandlerImpl{}
	cw.Append(h)

	tc := cw.Value()
	req.NotNil(tc)
	req.Equal(1, len(tc))
	req.Equal(h, tc[0])

	h2 := &testHandlerImpl2{}
	cw.Append(h2)

	tc = cw.Value()
	req.NotNil(tc)
	req.Equal(2, len(tc))
	req.Equal(h, tc[0])
	req.Equal(h2, tc[1])

	cw.Delete(h)
	tc = cw.Value()
	req.NotNil(tc)
	req.Equal(1, len(tc))
	req.Equal(h2, tc[0])

	cw.Delete(h)
	tc = cw.Value()
	req.NotNil(tc)
	req.Equal(1, len(tc))
	req.Equal(h2, tc[0])

	cw.Delete(h2)
	tc = cw.Value()
	req.Equal(0, len(tc))
}

func TestCloseSlice_Delete(t *testing.T) {
	t.Run("adding and removing a single element succeeds", func(t *testing.T) {
		req := require.New(t)
		cw := &CopyOnWriteSlice[testHandler]{}

		h := &testHandlerImpl{}
		cw.Append(h)
		cw.Delete(h)

		tc := cw.Value()
		req.NotNil(tc)
		req.Equal(0, len(tc))
	})

	t.Run("adding and removing a single element succeeds with DeleteIf", func(t *testing.T) {
		req := require.New(t)
		cw := &CopyOnWriteSlice[testHandler]{}

		h := &testHandlerImpl{}
		cw.Append(h)
		cw.DeleteIf(func(handler testHandler) bool {
			return handler == h
		})

		tc := cw.Value()
		req.NotNil(tc)
		req.Equal(0, len(tc))
	})

	t.Run("delete on empty cow slice does not panic", func(t *testing.T) {
		req := require.New(t)
		cw := &CopyOnWriteSlice[testHandler]{}

		h := &testHandlerImpl{}
		cw.Delete(h)

		tc := cw.Value()
		req.NotNil(tc)
		req.Equal(0, len(tc))
	})
}
