package cowslice

import (
	"reflect"
	"sync"
	"sync/atomic"
)

type CowSlice struct {
	value atomic.Value
	lock  sync.Mutex
}

func NewCowSlice(initialValue interface{}) *CowSlice {
	result := &CowSlice{}
	result.value.Store(initialValue)
	return result
}

func (slice *CowSlice) Value() interface{} {
	return slice.value.Load()
}

func Append(registry *CowSlice, listener interface{}) {
	registry.lock.Lock()
	defer registry.lock.Unlock()

	currentSlice := registry.value.Load()
	newSlice := reflect.Append(reflect.ValueOf(currentSlice), reflect.ValueOf(listener))
	registry.value.Store(newSlice.Interface())
}

func Delete(registry *CowSlice, listener interface{}) {
	registry.lock.Lock()
	defer registry.lock.Unlock()

	currentSlice := registry.value.Load()
	t := reflect.TypeOf(currentSlice)
	val := reflect.ValueOf(currentSlice)

	newSlice := reflect.MakeSlice(t, 0, val.Len()-1)
	for i := 0; i < val.Len(); i++ {
		next := val.Index(i)
		if next.Interface() != listener {
			newSlice = reflect.Append(newSlice, next)
		}
	}
	registry.value.Store(newSlice.Interface())
}
