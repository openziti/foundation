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
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/foundation/identity/identity"
	"github.com/openziti/foundation/transport"
	"io"
)

type classicListener struct {
	identity       *identity.TokenId
	endpoint       transport.Address
	socket         io.Closer
	close          chan struct{}
	handlers       []ConnectionHandler
	created        chan Underlay
	connectOptions ConnectOptions
	tcfg           transport.Configuration
	headers        map[int32][]byte
}

func NewClassicListener(identity *identity.TokenId, endpoint transport.Address, connectOptions ConnectOptions, headers map[int32][]byte) UnderlayListener {
	return NewClassicListenerWithTransportConfiguration(identity, endpoint, connectOptions, nil, headers)
}

func NewClassicListenerWithTransportConfiguration(identity *identity.TokenId, endpoint transport.Address, connectOptions ConnectOptions, tcfg transport.Configuration, headers map[int32][]byte) UnderlayListener {
	return &classicListener{
		identity:       identity,
		endpoint:       endpoint,
		close:          make(chan struct{}),
		created:        make(chan Underlay),
		connectOptions: connectOptions,
		tcfg:           tcfg,
		headers:        headers,
	}
}

func (listener *classicListener) Listen(handlers ...ConnectionHandler) error {
	incoming := make(chan transport.Connection, listener.connectOptions.MaxQueuedConnects)
	socket, err := listener.endpoint.Listen("classic", listener.identity, incoming, listener.tcfg)
	if err != nil {
		return err
	}
	listener.socket = socket
	listener.handlers = handlers

	for i := 0; i < listener.connectOptions.MaxOutstandingConnects; i++ {
		go listener.listener(incoming)
	}

	return nil
}

func (listener *classicListener) Close() error {
	close(listener.close)
	defer close(listener.created) // don't close created until socket is closed, to ensure we don't write to empty channel
	if err := listener.socket.Close(); err != nil {
		return err
	}
	listener.socket = nil
	return nil
}

func (listener *classicListener) Create(tcfg transport.Configuration) (Underlay, error) {
	listener.tcfg = tcfg
	if listener.created == nil {
		return nil, ListenerClosedError
	}
	select {
	case impl := <-listener.created:
		if impl != nil {
			return impl, nil
		}
	case <-listener.close:
	}
	return nil, ListenerClosedError
}

func (listener *classicListener) listener(incoming chan transport.Connection) {
	log := pfxlog.ContextLogger(listener.endpoint.String())
	log.Debug("started")
	defer log.Debug("exited")

	for {
		select {
		case peer := <-incoming:
			impl := newClassicImpl(peer, 2)
			if connectionId, err := globalRegistry.newConnectionId(); err == nil {
				impl.connectionId = connectionId

				if err := peer.SetReadTimeout(listener.connectOptions.ConnectTimeout()); err != nil {
					log.Errorf("could not set read timeout for [%s] (%v)", peer.Detail().Address, err)
					_ = peer.Close()
					return
				}

				request, hello, err := listener.receiveHello(impl)

				if err == nil {
					if err = peer.ClearReadTimeout(); err != nil {
						log.Errorf("could not clear read timeout for [%s] (%v)", peer.Detail().Address, err)
						_ = peer.Close()
						return
					}

					for _, h := range listener.handlers {
						if err := h.HandleConnection(hello, peer.PeerCertificates()); err != nil {
							log.Errorf("connection handler error for [%s] (%v)", peer.Detail().Address, err)
							if err := listener.ackHello(impl, request, false, err.Error()); err != nil {
								log.Errorf("error acknowledging hello for [%s] (%v)", peer.Detail().Address, err)
							}
							break
						}
					}

					impl.id = &identity.TokenId{Token: hello.IdToken}
					impl.headers = hello.Headers

					if err := listener.ackHello(impl, request, true, ""); err == nil {
						listener.created <- impl
					} else {
						log.Errorf("error acknowledging hello for [%s] (%v)", peer.Detail().Address, err)
					}

				} else {
					_ = peer.Close()
					log.Errorf("error receiving hello from [%s] (%v)", peer.Detail().Address, err)
				}
			} else {
				_ = peer.Close()
				log.Errorf("error getting connection id for [%s] (%v)", peer.Detail().Address, err)
			}

		case <-listener.close:
			return
		}
	}
}

func (listener *classicListener) receiveHello(impl *classicImpl) (*Message, *Hello, error) {
	log := pfxlog.ContextLogger(impl.Label())
	log.Debug("started")
	defer log.Debug("exited")

	request, err := impl.rxHello()
	if err != nil {
		if err == UnknownVersionError {
			writeUnknownVersionResponse(impl.peer.Writer())
		}
		_ = impl.Close()
		return nil, nil, fmt.Errorf("receive error (%s)", err)
	}
	if request.ContentType != ContentTypeHelloType {
		_ = impl.Close()
		return nil, nil, fmt.Errorf("unexpected content type [%d]", request.ContentType)
	}
	hello := UnmarshalHello(request)
	return request, hello, nil
}

func (listener *classicListener) ackHello(impl *classicImpl, request *Message, success bool, message string) error {
	response := NewResult(success, message)

	for key, val := range listener.headers {
		response.Headers[key] = val
	}

	response.Headers[ConnectionIdHeader] = []byte(impl.connectionId)
	response.sequence = HelloSequence

	response.ReplyTo(request)
	return impl.Tx(response)
}
