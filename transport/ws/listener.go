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

package ws

import (
	"crypto/tls"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/foundation/transport"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

var upgrader = websocket.Upgrader{}

type wsListener struct {
	log      *logrus.Entry
	incoming chan transport.Connection
	cfg      *WSConfig
}

/**
 *	Accept incoming HTTP connection, and upgrade it to a websocket suitable for comms between ziti-sdk-js and Ziti Edge Router
 */
func (listener *wsListener) handleWebsocket(w http.ResponseWriter, r *http.Request) {
	log := listener.log
	log.Info("entered")

	c, err := upgrader.Upgrade(w, r, nil) // upgrade from HTTP to binary socket

	if err != nil {
		log.WithField("err", err).Error("websocket upgrade failed. Failure not recoverable.")
	} else {

		connection := &Connection{
			detail: &transport.ConnectionDetail{
				Address: "ws:" + c.UnderlyingConn().RemoteAddr().String(),
				InBound: true,
				Name:    "ws",
			},
			ws:       c,
			log:      log,
			rxbuf:    newSafeBuffer(log),
			txbuf:    newSafeBuffer(log),
			tlsrxbuf: newSafeBuffer(log),
			tlstxbuf: newSafeBuffer(log),
			done:     make(chan struct{}),
			cfg:      listener.cfg,
			incoming: listener.incoming,
		}

		err := connection.tlsHandshake() // Do not proceed until the JS client can successfully complete a TLS handshake
		if err == nil {
			go connection.pinger()
			listener.incoming <- connection // pass the Websocket to the goroutine that will validate the HELLO handshake
		}
	}
}

func Listen(bindAddress string, name string, incoming chan transport.Connection, tcfg transport.Configuration) (io.Closer, error) {
	log := pfxlog.ContextLogger(name + "/ws:" + bindAddress)

	cfg := NewDefaultWSConfig()
	if tcfg != nil {
		if err := cfg.Load(tcfg); err != nil {
			return nil, errors.Wrap(err, "load configuration")
		}
	}
	logrus.Infof(cfg.Dump())

	go wslistener(log, bindAddress, cfg, name, incoming)

	return nil, nil
}

/**
 *	The TCP-based listener that accepts incoming HTTP connections that we will upgrade to Websocket connections.
 */
func wslistener(log *logrus.Entry, bindAddress string, cfg *WSConfig, name string, incoming chan transport.Connection) {

	log.Infof("starting HTTP (websocket) server at bindAddress [%s]", bindAddress)

	listener := &wsListener{
		log:      log,
		incoming: incoming,
		cfg:      cfg,
	}

	// Set up the HTTP -> Websocket upgrader options (once, before we start listening)
	upgrader.HandshakeTimeout = cfg.handshakeTimeout
	upgrader.ReadBufferSize = cfg.readBufferSize
	upgrader.WriteBufferSize = cfg.writeBufferSize
	upgrader.EnableCompression = cfg.enableCompression
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // Allow all origins

	router := mux.NewRouter()

	router.HandleFunc("/ws", listener.handleWebsocket).Methods("GET")

	cert, _ := tls.LoadX509KeyPair(cfg.serverCert, cfg.key)

	httpServer := &http.Server{
		Addr:         bindAddress,
		WriteTimeout: cfg.writeTimeout,
		ReadTimeout:  cfg.readTimeout,
		IdleTimeout:  cfg.idleTimeout,
		Handler:      router,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	if err := httpServer.ListenAndServeTLS("", ""); err != nil {
		panic(err)
	}
}
