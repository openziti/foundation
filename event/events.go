package event

import (
	"errors"
	"github.com/openziti/foundation/util/concurrenz"
)

type Dispatcher interface {
	Dispatch(event Event)
	Stop()
}

type Event interface {
	Handle()
}

var DispatcherNotRunningError = errors.New("dispatched not running")

func NewDispatcher() Dispatcher {
	result := &dispatcherImpl{
		shutdownC: make(chan struct{}),
		eventC:    make(chan Event, 25),
	}

	result.start()
	return result
}

type dispatcherImpl struct {
	running   concurrenz.AtomicBoolean
	shutdownC chan struct{}
	eventC    chan Event
}

func (dispatcher *dispatcherImpl) Stop() {
	if dispatcher.running.CompareAndSwap(true, false) {
		close(dispatcher.shutdownC)
	}
}

func (dispatcher *dispatcherImpl) start() {
	if dispatcher.running.CompareAndSwap(false, true) {
		go dispatcher.eventLoop()
	}
}

func (dispatcher *dispatcherImpl) eventLoop() {
	for dispatcher.running.Get() {
		select {
		case event := <-dispatcher.eventC:
			event.Handle()
		case <-dispatcher.shutdownC:
		}
	}
}

func (dispatcher *dispatcherImpl) Dispatch(event Event) {
	if dispatcher.running.Get() {
		select {
		case dispatcher.eventC <- event:
		case <-dispatcher.shutdownC:
		}
	}
}
