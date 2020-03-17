/*
	Copyright 2019 NetFoundry, Inc.

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

package tcp

import (
	"github.com/netfoundry/ziti-foundation/transport"
	"crypto/x509"
	"io"
	"net"
	"time"
)

type Connection struct {
	detail *transport.ConnectionDetail
	socket net.Conn
}

func (self *Connection) Detail() *transport.ConnectionDetail {
	return self.detail
}

func (self *Connection) PeerCertificates() []*x509.Certificate {
	return nil
}

func (self *Connection) Reader() io.Reader {
	return self.socket
}

func (self *Connection) ReadPeer(p []byte) (int, transport.Address, error) {
	n, err := self.socket.Read(p)
	return n, nil, err
}

func (self *Connection) Writer() io.Writer {
	return self.socket
}

func (self *Connection) Conn() net.Conn {
	return self.socket
}

func (self *Connection) SetReadTimeout(t time.Duration) error {
	return self.socket.SetReadDeadline(time.Now().Add(t))
}

func (self *Connection) SetWriteTimeout(t time.Duration) error {
	return self.socket.SetWriteDeadline(time.Now().Add(t))
}

func (self *Connection) Close() error {
	return self.socket.Close()
}
