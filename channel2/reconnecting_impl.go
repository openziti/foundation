/*
	Copyright 2019 Netfoundry, Inc.

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
	"github.com/netfoundry/ziti-foundation/identity/identity"
	"github.com/netfoundry/ziti-foundation/transport"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"sync"
)

func (impl *reconnectingImpl) Rx() (*Message, error) {
	log := pfxlog.ContextLogger(impl.Label())

	connected := true
	for {
		if connected {
			m, err := impl.rx()
			if err != nil {
				log.Errorf("rx error (%s). starting reconnection process", err)
				connected = false
			} else {
				return m, nil
			}
		} else {
			if err := impl.reconnectionHandler.Reconnect(impl); err != nil {
				log.Errorf("reconnection failed (%s)", err)
				return nil, fmt.Errorf("reconnection failed (%s)", err)
			} else {
				log.Info("reconnected")
				connected = true
			}
		}
	}
}

func (impl *reconnectingImpl) Tx(m *Message) error {
	log := pfxlog.ContextLogger(impl.Label())

	done := false
	connected := true
	for !done {
		if connected {
			if err := impl.tx(m); err != nil {
				log.Errorf("tx error (%s). starting reconnection process", err)
				connected = false
			} else {
				done = true
			}
		} else {
			if err := impl.reconnectionHandler.Reconnect(impl); err != nil {
				log.Errorf("reconnection failed (%s)", err)
				return fmt.Errorf("reconnection failed (%s)", err)

			} else {
				log.Info("reconnected")
				connected = true
			}
		}
	}
	return nil
}

func (impl *reconnectingImpl) Id() *identity.TokenId {
	return impl.id
}

func (impl *reconnectingImpl) Headers() map[int32][]byte {
	return impl.headers
}

func (impl *reconnectingImpl) LogicalName() string {
	return "reconnecting"
}

func (impl *reconnectingImpl) ConnectionId() string {
	return impl.connectionId
}

func (impl *reconnectingImpl) Certificates() []*x509.Certificate {
	return impl.peer.PeerCertificates()
}

func (impl *reconnectingImpl) Label() string {
	return fmt.Sprintf("u{%s}->i{%s}", impl.LogicalName(), impl.ConnectionId())
}

func (impl *reconnectingImpl) Close() error {
	impl.closeLock.Lock()
	defer impl.closeLock.Unlock()

	if !impl.closed {
		impl.closed = true
		return impl.peer.Close()
	}
	return nil
}

func (impl *reconnectingImpl) IsClosed() bool {
	return impl.closed
}

func newReconnectingImpl(peer transport.Connection, reconnectionHandler reconnectionHandler) *reconnectingImpl {
	return &reconnectingImpl{
		peer:                peer,
		reconnectionHandler: reconnectionHandler,
		readF:               readV2,
		marshalF:            marshalV2,
	}
}

func (impl *reconnectingImpl) setProtocolVersion(version uint32) {
	if version == 2 {
		impl.readF = readV2
		impl.marshalF = marshalV2
	} else {
		pfxlog.Logger().Warnf("asked to set unsupported protocol version %v", version)
	}
}

func (impl *reconnectingImpl) rx() (*Message, error) {
	if impl.closed {
		return nil, errors.New("underlay closed")
	}
	return impl.readF(impl.peer.Reader())
}

func (impl *reconnectingImpl) tx(m *Message) error {
	if impl.closed {
		return errors.New("underlay closed")
	}

	data, body, err := impl.marshalF(m)
	if err != nil {
		return err
	}

	_, err = impl.peer.Writer().Write(data)
	if err != nil {
		return err
	}

	_, err = impl.peer.Writer().Write(body)
	if err != nil {
		return err
	}

	return nil
}

// pingInstance currently does a single-sided (unverified) ping to see if the peer connection is functional.
//
func (impl *reconnectingImpl) pingInstance() error {
	log := pfxlog.ContextLogger(impl.Label())
	defer log.Info("exiting")
	log.Info("starting")

	ping := NewMessage(reconnectingPingContentType, nil)
	if err := impl.tx(ping); err != nil {
		return err
	}

	return nil
}

type reconnectingImpl struct {
	peer                transport.Connection
	id                  *identity.TokenId
	connectionId        string
	headers             map[int32][]byte
	reconnectionHandler reconnectionHandler
	closeLock           sync.Mutex
	closed              bool
	readF               readFunction
	marshalF            marshalFunction
}

type reconnectionHandler interface {
	Reconnect(impl *reconnectingImpl) error
}

const reconnectingPingContentType = -33
