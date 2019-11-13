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

package tls

import (
	"github.com/netfoundry/ziti-foundation/transport"
	"crypto/tls"
	"crypto/x509"
	"io"
	"net"
	"time"
)

type Connection struct {
	detail *transport.ConnectionDetail
	socket *tls.Conn
}

func (c *Connection) Detail() *transport.ConnectionDetail {
	return c.detail
}

func (c *Connection) PeerCertificates() []*x509.Certificate {
	return c.socket.ConnectionState().PeerCertificates
}

func (c *Connection) Reader() io.Reader {
	return c.socket
}

func (c *Connection) Writer() io.Writer {
	return c.socket
}

func (c *Connection) Conn() net.Conn {
	return c.socket
}

func (c *Connection) SetReadTimeout(t time.Duration) error {
	return c.socket.SetReadDeadline(time.Now().Add(t))
}

func (c *Connection) SetWriteTimeout(t time.Duration) error {
	return c.socket.SetWriteDeadline(time.Now().Add(t))
}

func (c *Connection) Close() error {
	return c.socket.Close()
}
