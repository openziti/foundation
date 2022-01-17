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

package channel

import (
	"context"
	"crypto/x509"
	"github.com/openziti/foundation/identity/identity"
	"github.com/openziti/foundation/transport"
	"io"
	"time"
)

// Channel represents an asyncronous, message-passing framework, designed to sit on top of an underlay.
//
type Channel interface {
	Identity
	SetLogicalName(logicalName string)
	Binding
	Sender
	io.Closer
	IsClosed() bool
	Underlay() Underlay
	StartRx()
	GetTimeSinceLastRead() time.Duration
}

type SendContext interface {
	Msg() *Message
	Priority() Priority
	WithPriority(p Priority) SendContext
	Context() context.Context
	WithTimeout(duration time.Duration) TimeoutSendContext
	Sequence() int32
	NotifyBeforeWrite()
	NotifyAfterWrite()
	NotifyErr(err error)
	ReplyChan() chan<- *Message
}

type TimeoutSendContext interface {
	SendContext
	SendAndWaitForWire(ch Channel) error
	SendForReply(ch Channel) (*Message, error)
	SendForTypedReply(ch Channel, result TypedMessage) error
}

type Identity interface {
	// The TokenId used to represent the identity of this channel to lower-level resources.
	//
	Id() *identity.TokenId

	// The LogicalName represents the purpose or usage of this channel (i.e. 'ctrl', 'mgmt' 'r/001', etc.) Usually used
	// by humans in understand the logical purpose of a channel.
	//
	LogicalName() string

	// The ConnectionId represents the identity of this Channel to internal API components ("instance identifier").
	// Usually used by the Channel framework to differentiate Channel instances.
	//
	ConnectionId() string

	// Certificates contains the identity certificates provided by the peer.
	//
	Certificates() []*x509.Certificate

	// Label constructs a consistently-formatted string used for context logging purposes, from the components above.
	//
	Label() string
}

// UnderlayListener represents a component designed to listen for incoming peer connections.
//
type UnderlayListener interface {
	Listen(handlers ...ConnectionHandler) error
	UnderlayFactory
	io.Closer
}

// UnderlayFactory is used by Channel to obtain an Underlay instance. An underlay "dialer" or "listener" implement
// UnderlayFactory, to provide instances to Channel.
//
type UnderlayFactory interface {
	Create(timeout time.Duration, tcfg transport.Configuration) (Underlay, error)
}

// Underlay abstracts a physical communications channel, typically sitting on top of 'transport'.
//
type Underlay interface {
	Rx() (*Message, error)
	Tx(m *Message) error
	Identity
	io.Closer
	IsClosed() bool
	Headers() map[int32][]byte
	SetWriteTimeout(duration time.Duration) error
}

type Sender interface {
	Send(sendCtx SendContext) error
}

const AnyContentType = -1
const HelloSequence = -1

var ListenerClosedError = listenerClosedError{}

type listenerClosedError struct{}

func (err listenerClosedError) Error() string {
	return "closed"
}
