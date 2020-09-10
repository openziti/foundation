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
	"bytes"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/foundation/identity/identity"
	"sync"
)

type wssImpl struct {
	peer         *WssConnection
	websocket    *websocket.Conn
	id           *identity.TokenId
	connectionId string
	headers      map[int32][]byte
	closeLock    sync.Mutex
	closed       bool
	// readF        readFunction
	marshalF marshalFunction
}

func (impl *wssImpl) rxHello() (*Message, error) {

	mt, wss_message, err := impl.websocket.ReadMessage()
	if err != nil {
		return nil, err
	}

	pfxlog.Logger().Infof("rxHello(): mt: %v, wss_message: %v", mt, wss_message)

	msg, _, _, err := readHello(bytes.NewReader([]byte(wss_message)))
	return msg, err
}

func (impl *wssImpl) Rx() (*Message, error) {
	if impl.closed {
		return nil, errors.New("underlay closed")
	}

	mt, wss_message, err := impl.websocket.ReadMessage()
	if err != nil {
		return nil, err
	}

	pfxlog.Logger().Infof("mt: %v, wss_message: %v", mt, wss_message)

	msg, err := readWssV2(wss_message)
	return msg, err
}

func (impl *wssImpl) Tx(m *Message) error {
	if impl.closed {
		return errors.New("underlay closed")
	}

	data, body, err := impl.marshalF(m)
	if err != nil {
		return err
	}

	pfxlog.Logger().Infof("data: %v, body: %v", data, body)

	err = impl.websocket.WriteMessage(websocket.BinaryMessage, append(data[:], body[:]...))
	if err != nil {
		pfxlog.Logger().Errorf("Tx(): err: %v", err)
		return err
	}

	return nil
}

func (impl *wssImpl) Id() *identity.TokenId {
	return impl.id
}

func (impl *wssImpl) Headers() map[int32][]byte {
	return impl.headers
}

func (impl *wssImpl) LogicalName() string {
	return "wss"
}

func (impl *wssImpl) ConnectionId() string {
	return impl.connectionId
}

func (impl *wssImpl) Certificates() []*x509.Certificate {
	return impl.peer.PeerCertificates()
}

func (impl *wssImpl) Label() string {
	return fmt.Sprintf("u{%s}->i{%s}", impl.LogicalName(), impl.ConnectionId())
}

func (impl *wssImpl) Close() error {
	impl.closeLock.Lock()
	defer impl.closeLock.Unlock()

	if !impl.closed {
		impl.closed = true
		return impl.peer.Close()
	}
	return nil
}

func (impl *wssImpl) IsClosed() bool {
	return impl.closed
}

func newWssImpl(peer *WssConnection, version uint32) *wssImpl {
	marshalF := marshalV2
	return &wssImpl{
		peer:      peer,
		websocket: peer.Websocket(),
		// readF:     readF,
		marshalF: marshalF,
	}
}
