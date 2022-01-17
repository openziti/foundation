package channel

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/michaelquigley/pfxlog"
	"github.com/pkg/errors"
	"reflect"
	"time"
)

type MutableSendContext struct {
	*Message
	p       Priority
	context context.Context
	cancelF context.CancelFunc
}

func (self *MutableSendContext) Context() context.Context {
	return self.context
}

func (self *MutableSendContext) WithTimeout(duration time.Duration) TimeoutSendContext {
	parent := self.context
	if parent == nil {
		parent = context.Background()
	}
	self.context, self.cancelF = context.WithTimeout(parent, duration)
	return self
}

func (self *MutableSendContext) Priority() Priority {
	return self.p
}

func (self *MutableSendContext) WithPriority(p Priority) SendContext {
	self.p = p
	return self
}

func (self *MutableSendContext) SendAndWaitForWire(ch Channel) error {
	waitSendContext := &WaitSendContext{MutableSendContext: self}
	return waitSendContext.WaitForWire(ch)
}

func (self *MutableSendContext) SendForReply(ch Channel) (*Message, error) {
	replyContext := &ReplySendContext{MutableSendContext: self}
	return replyContext.WaitForReply(ch)
}

func (self *MutableSendContext) SendForTypedReply(ch Channel, result TypedMessage) error {
	replyContext := &ReplySendContext{MutableSendContext: self}
	return replyContext.WaitForTypedReply(ch, result)
}

type WaitSendContext struct {
	*MutableSendContext
	errC chan error
}

func (self *WaitSendContext) NotifyAfterWrite() {
	close(self.errC)
}

func (self *WaitSendContext) NotifyErr(err error) {
	self.errC <- err
}

func (self *WaitSendContext) WaitForWire(ch Channel) error {
	if err := self.context.Err(); err != nil {
		return err
	}

	defer self.cancelF()

	self.errC = make(chan error, 1)

	if err := ch.Send(self); err != nil {
		return err
	}
	select {
	case err := <-self.errC:
		return err
	case <-self.context.Done():
		return errors.New("timeout waiting for message to be written to wire")
	}
}

type ReplySendContext struct {
	*MutableSendContext
	errC   chan error
	replyC chan *Message
}

func (self *ReplySendContext) NotifyErr(err error) {
	self.errC <- err
}

func (self *ReplySendContext) ReplyChan() chan<- *Message {
	return self.replyC
}

func (self *ReplySendContext) WaitForReply(ch Channel) (*Message, error) {
	if err := self.context.Err(); err != nil {
		return nil, err
	}

	defer self.cancelF()

	self.errC = make(chan error, 1)
	self.replyC = make(chan *Message, 1)

	if err := ch.Send(self); err != nil {
		return nil, err
	}

	select {
	case err := <-self.errC:
		return nil, err
	case <-self.context.Done():
		return nil, errors.New("timeout waiting for message to be written to wire")
	case reply := <-self.replyC:
		return reply, nil
	}
}

func (self *ReplySendContext) WaitForTypedReply(ch Channel, result TypedMessage) error {
	responseMsg, err := self.WaitForReply(ch)
	if err != nil {
		return err
	}
	if responseMsg.ContentType != result.GetContentType() {
		return errors.Errorf("unexpected response type %v to request of type %v. expected %v",
			responseMsg.ContentType, self.Msg().ContentType, result.GetContentType())
	}
	if err = proto.Unmarshal(responseMsg.Body, result); err != nil {
		return err
	}
	return nil
}

type ErrorSendContext struct {
	ctx context.Context
}

func (self *ErrorSendContext) WithTimeout(duration time.Duration) TimeoutSendContext {
	panic("implement me")
}

func (self *ErrorSendContext) Msg() *Message {
	return nil
}

func (self *ErrorSendContext) Priority() Priority {
	return Standard
}

func (self *ErrorSendContext) WithPriority(Priority) SendContext {
	return self
}

func (self *ErrorSendContext) Context() context.Context {
	return self.ctx
}

func (self *ErrorSendContext) Sequence() int32 {
	return 0
}

func (self *ErrorSendContext) NotifyBeforeWrite() {}

func (self *ErrorSendContext) NotifyAfterWrite() {}

func (self *ErrorSendContext) NotifyErr(error) {}

func (self *ErrorSendContext) ReplyChan() chan<- *Message {
	return nil
}

func MarshalProto(contentType int32, msg proto.Message) SendContext {
	b, err := proto.Marshal(msg)
	if err != nil {
		return &ErrorSendContext{
			ctx: NewErrorContext(errors.Wrapf(err, "failed to marshal %v", reflect.TypeOf(msg))),
		}
	}
	return NewMessage(contentType, b)
}

func MarshalTyped(msg TypedMessage) SendContext {
	return MarshalProto(msg.GetContentType(), msg)
}

func NewErrorContext(err error) context.Context {
	result := &ErrorContext{
		err:     err,
		closedC: make(chan struct{}),
	}
	close(result.closedC)
	return result
}

type ErrorContext struct {
	err     error
	closedC chan struct{}
}

func (self *ErrorContext) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

func (self *ErrorContext) Done() <-chan struct{} {
	return self.closedC
}

func (self *ErrorContext) Err() error {
	return self.err
}

func (self *ErrorContext) Value(key interface{}) interface{} {
	// ignore for now. may need an implementation at some point
	pfxlog.Logger().Error("ErrorContext.Value called, but not implemented!!!")
	return nil
}
