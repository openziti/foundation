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
	"errors"
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/netfoundry/ziti-foundation/identity/identity"
	"github.com/netfoundry/ziti-foundation/transport"
	"io"
	"time"
)

const (
	helloTimeout = 3 * time.Second
)

type classicListener struct {
	identity *identity.TokenId
	endpoint transport.Address
	socket   io.Closer
	close    chan struct{}
	handlers []ConnectionHandler
	created  chan Underlay
}

func NewClassicListener(identity *identity.TokenId, endpoint transport.Address) UnderlayListener {
	return &classicListener{
		identity: identity,
		endpoint: endpoint,
		close:    make(chan struct{}),
		created:  make(chan Underlay),
	}
}

func (listener *classicListener) Listen(handlers ...ConnectionHandler) error {
	incoming := make(chan transport.Connection)
	socket, err := listener.endpoint.Listen("classic", listener.identity, incoming)
	if err != nil {
		return err
	}
	listener.socket = socket
	listener.handlers = handlers
	go listener.listener(incoming)

	return nil
}

func (listener *classicListener) Close() error {
	close(listener.close)
	close(listener.created)
	if err := listener.socket.Close(); err != nil {
		return err
	}
	listener.socket = nil
	return nil
}

func (listener *classicListener) Create() (Underlay, error) {
	if listener.created == nil {
		return nil, errors.New("closed")
	}
	impl := <-listener.created
	if impl == nil {
		return nil, errors.New("closed")
	}
	return impl, nil
}

func (listener *classicListener) listener(incoming chan transport.Connection) {
	logger := pfxlog.ContextLogger(listener.endpoint.String())
	logger.Debug("listener waiting for connections: started")
	defer logger.Debug("listener stopped waiting for connections: exited")

	for {
		select {
		case peer := <-incoming:
			go listener.onConnect(peer)
		case <-listener.close:
			return
		}
	}
}

func (listener *classicListener) onConnect(peer transport.Connection) {
	logger := pfxlog.ContextLogger(listener.endpoint.String())
	impl := newClassicImpl(peer, 2)
	if connectionId, err := globalRegistry.newConnectionId(); err == nil {
		impl.connectionId = connectionId
		request, hello, err := listener.receiveHello(impl)
		if err == nil {
			for _, h := range listener.handlers {
				logger.Infof("hello: %v, peer: %v, handler: %v", hello, peer, h)
				if err := h.HandleConnection(hello, peer.PeerCertificates()); err != nil {
					logger.Errorf("connection handler error (%s)", err)
					if err := listener.ackHello(impl, request, false, err.Error()); err != nil {
						logger.Errorf("error acknowledging hello (%s)", err)
					}
					break
				}
			}

			impl.id = &identity.TokenId{Token: hello.IdToken}
			impl.headers = hello.Headers

			if err := listener.ackHello(impl, request, true, ""); err == nil {
				listener.created <- impl
			} else {
				logger.Errorf("error acknowledging hello (%s)", err)
			}

		} else {
			logger.Errorf("error receiving hello (%s)", err)
		}
	} else {
		logger.Errorf("error getting connection id (%s)", err)
	}

}

type helloResponse struct {
	message *Message
	hello   *Hello
	error   error
}

func (listener *classicListener) receiveHello(impl *classicImpl) (*Message, *Hello, error) {
	responseChan := make(chan helloResponse)

	go func() {
		log := pfxlog.ContextLogger(impl.Label())
		log.Debug("started")
		defer log.Debug("exited")

		request, err := impl.rxHello()
		if err != nil {
			if err == UnknownVersionError {
				writeUnknownVersionResponse(impl.peer.Writer())
			}
			_ = impl.Close()
			resp := helloResponse{error: fmt.Errorf("receive error (%s)", err)}
			responseChan <- resp
		}
		if request.ContentType != ContentTypeHelloType {
			_ = impl.Close()
			resp := helloResponse{error: fmt.Errorf("unexpected content type [%d]", request.ContentType)}
			responseChan <- resp
		}
		hello := UnmarshalHello(request)
		responseChan <- helloResponse{message: request, hello: hello}
	}()

	select {
	case helloResp := <-responseChan:

		return helloResp.message, helloResp.hello, helloResp.error
	case <-time.After(helloTimeout):
		if err := impl.Close(); err != nil {
			return nil, nil, fmt.Errorf("hello timed out, closing errored: %s", err)
		}
		return nil, nil, fmt.Errorf("hello timed out")
	}
}

func (listener *classicListener) ackHello(impl *classicImpl, request *Message, success bool, message string) error {
	response := NewResult(success, message)
	response.Headers[ConnectionIdHeader] = []byte(impl.connectionId)
	response.sequence = HelloSequence
	response.ReplyTo(request)
	return impl.Tx(response)
}
