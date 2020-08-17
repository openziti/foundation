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
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/foundation/identity/identity"
	"github.com/openziti/foundation/transport"
	"io"
	"net"
	"net/http"
	"time"
)

type wssListener struct {
	identity       *identity.TokenId
	endpoint       transport.Address
	socket         io.Closer
	close          chan struct{}
	handlers       []ConnectionHandler
	created        chan Underlay
	connectOptions ConnectOptions
	serverCert     string
	key            string
	tcfg           transport.Configuration
	incoming       chan *WssConnection
}

func NewWssListener(identity *identity.TokenId, endpoint transport.Address, connectOptions ConnectOptions, serverCert string, key string) UnderlayListener {
	return NewWssListenerWithTransportConfiguration(identity, endpoint, connectOptions, nil, serverCert, key)
}

func NewWssListenerWithTransportConfiguration(identity *identity.TokenId, endpoint transport.Address, connectOptions ConnectOptions, tcfg transport.Configuration, serverCert string, key string) UnderlayListener {
	return &wssListener{
		identity:       identity,
		endpoint:       endpoint,
		close:          make(chan struct{}),
		created:        make(chan Underlay),
		incoming:       make(chan *WssConnection, connectOptions.MaxQueuedConnects),
		connectOptions: connectOptions,
		serverCert:     serverCert,
		key:            key,
		tcfg:           tcfg,
	}
}

type WssConnection struct {
	detail    *transport.ConnectionDetail
	websocket *websocket.Conn
	socket    *tls.Conn
}

func (c *WssConnection) Detail() *transport.ConnectionDetail {
	return c.detail
}

func (c *WssConnection) PeerCertificates() []*x509.Certificate {
	return c.socket.ConnectionState().PeerCertificates
}

func (c *WssConnection) Reader() io.Reader {
	return c.socket
}

func (c *WssConnection) Writer() io.Writer {
	return c.socket
}

func (c *WssConnection) Websocket() *websocket.Conn {
	return c.websocket
}

func (c *WssConnection) Conn() net.Conn {
	return c.socket
}

func (c *WssConnection) SetReadTimeout(t time.Duration) error {
	return c.socket.SetReadDeadline(time.Now().Add(t))
}

func (c *WssConnection) SetWriteTimeout(t time.Duration) error {
	return c.socket.SetWriteDeadline(time.Now().Add(t))
}

func (c *WssConnection) Close() error {
	return c.socket.Close()
}

var upgrader = websocket.Upgrader{} // use default options

/**
 *	Accept incoming HTTP connection, and upgrade it to a websocket suitable for comms between browZer and Ziti edge router
 */
func (listener *wssListener) handleWebsocket(w http.ResponseWriter, r *http.Request) {
	log := pfxlog.ContextLogger(listener.endpoint.String())
	log.Info("entered")

	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // Allow all origins

	c, err := upgrader.Upgrade(w, r, nil) // upgrade from HTTP to binary socket

	if err != nil {
		log.WithField("err", err).Error("websocket upgrade failed. Failure not recoverable. Exiting listen loop")
	} else {

		socket := c.UnderlyingConn() // Obtain the socket underneath the websocket

		connection := &WssConnection{
			detail: &transport.ConnectionDetail{
				Address: "wss:" + socket.RemoteAddr().String(),
				InBound: true,
				Name:    "wss",
			},
			websocket: c,
			socket:    socket.(*tls.Conn),
		}

		listener.incoming <- connection // pass the socket to the goroutine that will validate the HELLO handshake
	}
}

/**
 *	Simply start the goroutines that will process incoming websocket connections
 */
func (listener *wssListener) Listen(handlers ...ConnectionHandler) error {
	go listener.wsslistener()
	go listener.binaryListener()
	return nil
}

/**
 *	The TLS-based listener that accepts incoming HTTP connections that we need to upgrade to websocket connections
 */
func (listener *wssListener) wsslistener() {
	log := pfxlog.ContextLogger(listener.endpoint.String())
	log.Info("starting HTTP (websocket) server")

	router := mux.NewRouter()

	router.HandleFunc("/wss", listener.handleWebsocket).Methods("GET")

	httpServer := &http.Server{
		Addr:         listener.endpoint.BindableAddress(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequestClientCert,
		},
	}

	// log.Info("listener.options: [%v]", listener.options)

	if err := httpServer.ListenAndServeTLS(listener.serverCert, listener.key); err != nil {
		panic(err)
	}
}

func (listener *wssListener) Close() error {
	close(listener.close)
	close(listener.created)
	if err := listener.socket.Close(); err != nil {
		return err
	}
	listener.socket = nil
	return nil
}

func (listener *wssListener) Create(tcfg transport.Configuration) (Underlay, error) {
	if listener.created == nil {
		return nil, errors.New("closed")
	}
	impl := <-listener.created
	if impl == nil {
		return nil, errors.New("closed")
	}
	return impl, nil
}

/**
 *	Take the incoming (now binary) websocket
 */
func (listener *wssListener) binaryListener() {
	log := pfxlog.ContextLogger(listener.endpoint.String())
	log.Info("started")
	defer log.Info("exited")

	for {
		select {
		case peer := <-listener.incoming:

			impl := newWssImpl(peer, 2)
			if connectionId, err := globalRegistry.newConnectionId(); err == nil {
				log.Infof("hello: new connectionId: %v", connectionId)

				impl.connectionId = connectionId
				request, hello, err := listener.receiveHello(impl)
				if err == nil {
					for _, h := range listener.handlers {
						log.Infof("hello: %v, peer: %v, handler: %v", hello, peer, h)
						if err := h.HandleConnection(hello, peer.PeerCertificates()); err != nil {
							log.Errorf("connection handler error (%s)", err)
							if err := listener.ackHello(impl, request, false, err.Error()); err != nil {
								log.Errorf("error acknowledging hello (%s)", err)
							}
							break
						}
					}

					impl.id = &identity.TokenId{Token: hello.IdToken}
					impl.headers = hello.Headers

					if err := listener.ackHello(impl, request, true, ""); err == nil {
						listener.created <- impl
					} else {
						log.Errorf("error acknowledging hello (%s)", err)
					}

				} else {
					log.Errorf("error receiving hello (%s)", err)
				}
			} else {
				log.Errorf("error getting connection id (%s)", err)
			}

		case <-listener.close:
			return
		}
	}
}

func (listener *wssListener) receiveHello(impl *wssImpl) (*Message, *Hello, error) {
	log := pfxlog.ContextLogger(impl.Label())
	log.Debug("started")
	defer log.Debug("exited")

	request, err := impl.rxHello()
	if err != nil {
		if err == UnknownVersionError {
			// writeUnknownVersionResponse(impl.peer.Writer())
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

func (listener *wssListener) ackHello(impl *wssImpl, request *Message, success bool, message string) error {
	response := NewResult(success, message)
	response.Headers[ConnectionIdHeader] = []byte(impl.connectionId)
	response.sequence = HelloSequence
	response.ReplyTo(request)
	return impl.Tx(response)
}
