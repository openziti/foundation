package concurrenz

import "sync/atomic"

type RefCount int32

func (self *RefCount) IncrRefCount() int32 {
	return atomic.AddInt32((*int32)(self), 1)
}

func (self *RefCount) DecrRefCount() int32 {
	return atomic.AddInt32((*int32)(self), -1)
}
