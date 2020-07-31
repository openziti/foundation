/*
	Copyright NetFoundry, Inc.

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

package channel2

import (
	"container/heap"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/foundation/identity/identity"
	"github.com/openziti/foundation/transport"
	"github.com/openziti/foundation/util/info"
	"github.com/openziti/foundation/util/sequence"
	"io"
	"sync"
	"time"
)

type channelImpl struct {
	logicalName       string
	underlayFactory   UnderlayFactory
	underlay          Underlay
	options           *Options
	sequence          *sequence.Sequence
	outQueue          chan *priorityMessage
	outPriority       *priorityHeap
	waiters           sync.Map
	syncers           sync.Map
	closed            bool
	closeLock         sync.Mutex
	peekHandlers      []PeekHandler
	transformHandlers []TransformHandler
	receiveHandlers   map[int32]ReceiveHandler
	errorHandlers     []ErrorHandler
	closeHandlers     []CloseHandler
	userData          interface{}
	lastActivity      int64
}

func NewChannel(logicalName string, underlayFactory UnderlayFactory, options *Options) (Channel, error) {
	return NewChannelWithTransportConfig(logicalName, underlayFactory, options, nil)
}

func NewChannelWithTransportConfig(logicalName string, underlayFactory UnderlayFactory, options *Options, tcfg transport.Configuration) (Channel, error) {
	impl := &channelImpl{
		logicalName:     logicalName,
		underlayFactory: underlayFactory,
		options:         options,
		sequence:        sequence.NewSequence(),
		outQueue:        make(chan *priorityMessage, 4),
		outPriority:     &priorityHeap{},
		receiveHandlers: make(map[int32]ReceiveHandler),
	}

	heap.Init(impl.outPriority)
	impl.AddReceiveHandler(&pingHandler{})

	underlay, err := underlayFactory.Create(tcfg)
	if err != nil {
		return nil, err
	}
	impl.underlay = underlay

	if options != nil {
		for _, binder := range options.BindHandlers {
			if err := binder.BindChannel(impl); err != nil {
				return nil, err
			}
		}
	}

	impl.startMultiplex()
	//go impl.keepalive()

	return impl, nil
}

func (channel *channelImpl) Id() *identity.TokenId {
	return channel.underlay.Id()
}

func (channel *channelImpl) LogicalName() string {
	return channel.logicalName
}

func (channel *channelImpl) SetLogicalName(logicalName string) {
	channel.logicalName = logicalName
}

func (channel *channelImpl) ConnectionId() string {
	return channel.underlay.ConnectionId()
}

func (channel *channelImpl) Certificates() []*x509.Certificate {
	return channel.underlay.Certificates()
}

func (channel *channelImpl) Label() string {
	if channel.underlay != nil {
		return fmt.Sprintf("ch{%s}->%s", channel.LogicalName(), channel.underlay.Label())
	} else {
		return fmt.Sprintf("ch{%s}->{}", channel.LogicalName())
	}
}

func (channel *channelImpl) Bind(h BindHandler) error {
	return h.BindChannel(channel)
}

func (channel *channelImpl) AddPeekHandler(h PeekHandler) {
	channel.peekHandlers = append(channel.peekHandlers, h)
}

func (channel *channelImpl) AddTransformHandler(h TransformHandler) {
	channel.transformHandlers = append(channel.transformHandlers, h)
}

func (channel *channelImpl) AddReceiveHandler(h ReceiveHandler) {
	channel.receiveHandlers[h.ContentType()] = h
}

func (channel *channelImpl) AddErrorHandler(h ErrorHandler) {
	channel.errorHandlers = append(channel.errorHandlers, h)
}

func (channel *channelImpl) AddCloseHandler(h CloseHandler) {
	channel.closeHandlers = append(channel.closeHandlers, h)
}

func (channel *channelImpl) SetUserData(data interface{}) {
	channel.userData = data
}

func (channel *channelImpl) GetUserData() interface{} {
	return channel.userData
}

func (channel *channelImpl) Close() error {
	channel.closeLock.Lock()
	defer channel.closeLock.Unlock()

	if !channel.closed {
		pfxlog.ContextLogger(channel.Label()).Debug("closing channel")

		channel.closed = true

		close(channel.outQueue)

		for _, peekHandler := range channel.peekHandlers {
			peekHandler.Close(channel)
		}

		if len(channel.closeHandlers) > 0 {
			for _, closeHandler := range channel.closeHandlers {
				closeHandler.HandleClose(channel)
			}
		} else {
			pfxlog.ContextLogger(channel.Label()).Debug("no close handlers")
		}

		return channel.underlay.Close()
	}

	return nil
}

func (channel *channelImpl) IsClosed() bool {
	return channel.closed
}

func (channel *channelImpl) Send(m *Message) error {
	return channel.SendWithPriority(m, Standard)
}

func (channel *channelImpl) SendWithPriority(m *Message, p Priority) (err error) {
	if channel.closed {
		return errors.New("channel closed")
	}

	defer func() {
		if r := recover(); r != nil {
			pfxlog.ContextLogger(channel.Label()).Error("send on closed channel")
			err = errors.New("send on closed channel")
			return
		}
	}()
	channel.stampSequence(m)
	channel.outQueue <- &priorityMessage{m: m, p: p}
	return nil
}

func (channel *channelImpl) SendAndSync(m *Message) (chan error, error) {
	return channel.SendAndSyncWithPriority(m, Standard)
}

func (channel *channelImpl) SendAndSyncWithPriority(m *Message, p Priority) (syncCh chan error, err error) {
	if channel.closed {
		return nil, errors.New("channel closed")
	}

	defer func() {
		if r := recover(); r != nil {
			pfxlog.ContextLogger(channel.Label()).Error("send on closed channel")
			err = errors.New("send on closed channel")
			return
		}
	}()
	channel.stampSequence(m)
	syncCh = make(chan error, 1)
	channel.syncers.Store(m.sequence, syncCh)
	channel.outQueue <- &priorityMessage{m: m, p: p}
	return syncCh, nil
}

func (channel *channelImpl) SendWithTimeout(m *Message, timeout time.Duration) error {
	syncC, err := channel.SendAndSync(m)
	if err != nil {
		return err
	}
	select {
	case err := <-syncC:
		return err
	case <-time.After(timeout):
		return errors.New("write deadline exceeded")
	}
}

func (channel *channelImpl) SendAndWaitWithTimeout(m *Message, timeout time.Duration) (*Message, error) {
	replyChan, err := channel.SendAndWait(m)
	if err != nil {
		return nil, err
	}
	select {
	case replyMsg := <-replyChan:
		return replyMsg, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout waiting for response")
	}
}

func (channel *channelImpl) SendAndWait(m *Message) (chan *Message, error) {
	return channel.SendAndWaitWithPriority(m, Standard)
}

func (channel *channelImpl) SendAndWaitWithPriority(m *Message, p Priority) (waitCh chan *Message, err error) {
	if channel.closed {
		return nil, errors.New("channel closed")
	}

	defer func() {
		if r := recover(); r != nil {
			pfxlog.ContextLogger(channel.Label()).Error("send on closed channel")
			err = errors.New("send on closed channel")
			return
		}
	}()
	channel.stampSequence(m)
	waitCh = make(chan *Message, 1)
	channel.waiters.Store(m.sequence, waitCh)
	channel.outQueue <- &priorityMessage{m: m, p: p}
	return waitCh, nil
}

func (channel *channelImpl) SendForReply(msg TypedMessage, timeout time.Duration) (*Message, error) {
	body, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	envelopeMsg := NewMessage(msg.GetContentType(), body)
	waitCh, err := channel.SendAndWait(envelopeMsg)
	if err != nil {
		return nil, err
	}

	select {
	case responseMsg := <-waitCh:
		return responseMsg, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("timed out waiting for response to request of type %v", msg.GetContentType())
	}
}

func (channel *channelImpl) SendForReplyAndDecode(msg TypedMessage, timeout time.Duration, result TypedMessage) error {
	responseMsg, err := channel.SendForReply(msg, timeout)
	if err != nil {
		return err
	}
	if responseMsg.ContentType != result.GetContentType() {
		return fmt.Errorf("unexpected response type %v to request of type %v. expected %v",
			responseMsg.ContentType, msg.GetContentType(), result.GetContentType())
	}
	if err := proto.Unmarshal(responseMsg.Body, result); err != nil {
		return err
	}
	return nil
}

func (channel *channelImpl) Underlay() Underlay {
	return channel.underlay
}

func (channel *channelImpl) startMultiplex() {
	for _, peekHandler := range channel.peekHandlers {
		peekHandler.Connect(channel, "")
	}

	go channel.rxer()
	go channel.txer()
}

func (channel *channelImpl) rxer() {
	log := pfxlog.ContextLogger(channel.Label())
	log.Debug("started")
	defer log.Debug("exited")

	defer func() { _ = channel.Close() }()
	defer func() {
		channel.waiters.Range(func(k, v interface{}) bool {
			channel.waiters.Delete(k)
			return true
		})
	}()

	for {
		m, err := channel.underlay.Rx()
		if err != nil {
			if err == io.EOF {
				log.Debug("EOF")
			} else {
				log.Errorf("rx error (%s)", err)
			}
			return
		}
		channel.lastActivity = info.NowInMilliseconds()

		for _, transformHandler := range channel.transformHandlers {
			transformHandler.Rx(m, channel)
		}

		for _, peekHandler := range channel.peekHandlers {
			peekHandler.Rx(m, channel)
		}

		handled := false
		if m.IsReply() {
			replyFor := m.ReplyFor()
			if tmp, found := channel.waiters.Load(replyFor); found {
				log.Debugf("waiter found for message. type [%v], sequence [%v], replyFor [%v]", m.ContentType, m.sequence, replyFor)

				waiter := tmp.(chan *Message)
				select {
				case waiter <- m:
				default:
					log.Warnf("unable to notify waiter of response. type [%v], sequence [%v], replyFor [%v]", m.ContentType, m.sequence, replyFor)
				}
				channel.waiters.Delete(replyFor)
				handled = true
			} else {
				log.Warnf("no waiter for for message. type [%v], sequence [%v], replyFor [%v]", m.ContentType, m.sequence, replyFor)
			}
		}

		if !handled {
			if receiveHandler, found := channel.receiveHandlers[m.ContentType]; found {
				receiveHandler.HandleReceive(m, channel)

			} else if anyHandler, found := channel.receiveHandlers[AnyContentType]; found {
				anyHandler.HandleReceive(m, channel)

			} else {
				log.Warnf("dropped message [%d]", m.ContentType)
			}
		}
	}
}

func (channel *channelImpl) txer() {
	log := pfxlog.ContextLogger(channel.Label())
	defer log.Debug("exited")
	log.Debug("started")

	for {
		done := false
		selecting := true

		pm, ok := <-channel.outQueue
		if ok {
			heap.Push(channel.outPriority, pm)
		} else {
			done = true
			selecting = false
		}

		for selecting {
			select {
			case pm, ok := <-channel.outQueue:
				if ok {
					heap.Push(channel.outPriority, pm)
				} else {
					done = true
					selecting = false
				}
			default:
				selecting = false
			}
		}

		for channel.outPriority.Len() > 0 {
			pm := heap.Pop(channel.outPriority).(*priorityMessage)
			m := pm.m

			for _, transformHandler := range channel.transformHandlers {
				transformHandler.Tx(m, channel)
			}

			syncCh := channel.getSyncer(m)
			var syncErr error = nil
			if err := channel.underlay.Tx(m); err != nil {
				log.Errorf("error tx (%s)", err)
				syncErr = err
				done = true
			}
			channel.lastActivity = info.NowInMilliseconds()

			for _, peekHandler := range channel.peekHandlers {
				peekHandler.Tx(m, channel)
			}

			if syncCh != nil {
				select {
				case syncCh <- syncErr:
				default:
					log.Warn("unable to notify syncer")
				}

			}
			if syncErr != nil {
				for _, errorHandler := range channel.errorHandlers {
					errorHandler.HandleError(syncErr, channel)
				}
			}
		}

		if done {
			return
		}
	}
}

func (channel *channelImpl) stampSequence(m *Message) {
	m.sequence = int32(channel.sequence.Next())
}

func (channel *channelImpl) getSyncer(m *Message) chan error {
	if syncCh, found := channel.syncers.Load(m.sequence); found {
		channel.syncers.Delete(m.sequence)
		return syncCh.(chan error)
	}
	return nil
}

func (channel *channelImpl) keepalive() {
	log := pfxlog.ContextLogger(channel.Label())
	log.Debug("started")
	defer log.Debug("exited")
	defer func() { _ = channel.Close() }()

	for {
		time.Sleep(1 * time.Second)
		if channel.IsClosed() {
			return
		}

		now := info.NowInMilliseconds()
		if now-channel.lastActivity > 15000 {
			request := NewMessage(ContentTypePingType, nil)
			waitCh, err := channel.SendAndWaitWithPriority(request, High)
			if err == nil {
				select {
				case response := <-waitCh:
					if response != nil {
						if response.ContentType == ContentTypeResultType {
							result := UnmarshalResult(response)
							if !result.Success {
								log.Error("failed ping response")
							} else {
								log.Debug("ping success")
							}
						} else {
							log.Errorf("unexpected ping response [%d]", response.ContentType)
						}
					} else {
						log.Error("wait channel closed")
						return
					}
				case <-time.After(time.Millisecond * 10000):
					log.Error("ping timeout")
					return
				}
			} else {
				log.Errorf("unexpected error (%s)", err)
			}
		}
	}
}
