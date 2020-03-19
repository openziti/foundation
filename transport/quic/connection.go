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

package quic

import (
	"github.com/netfoundry/ziti-foundation/transport"
	"crypto/x509"
	"fmt"
	"github.com/lucas-clemente/quic-go"
	"io"
	"net"
	"time"
)

// Connection represents a concrete QUIC connection.
//
type Connection struct {
	detail  *transport.ConnectionDetail
	session quic.Session
	stream  quic.Stream
}

func (c *Connection) Detail() *transport.ConnectionDetail {
	return c.detail
}

func (c *Connection) PeerCertificates() []*x509.Certificate {
	return c.session.ConnectionState().PeerCertificates
}

// Reader method on the transport.Connection interface.
//
func (c *Connection) Reader() io.Reader {
	return c.stream
}

// Writer method on the transport.Connection interface.
//
func (c *Connection) Writer() io.Writer {
	return c.stream
}

// QUIC doesn't provide an underlying net.Conn
//
func (c *Connection) Conn() net.Conn {
	return nil
}

// SetReadTimeout sets the read timeout to the provided duration.
//
func (c *Connection) SetReadTimeout(t time.Duration) error {
	return c.stream.SetReadDeadline(time.Now().Add(t))
}

// SetWriteTimeout sets the write timeout to the provided duration.
//
func (c *Connection) SetWriteTimeout(t time.Duration) error {
	return c.stream.SetWriteDeadline(time.Now().Add(t))
}

// Close method on the transport.Connection interface.
//
func (c *Connection) Close() error {
	err1 := c.stream.Close()
	err2 := c.session.Close()

	if err1 == nil {
		return err2
	}
	if err2 == nil {
		return err1
	}
	return fmt.Errorf("multiple errors closing quic connection (%s), (%s)", err1, err2)
}
