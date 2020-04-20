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
	"crypto/x509"
	"fmt"
	"github.com/lucas-clemente/quic-go"
	"github.com/netfoundry/ziti-foundation/transport"
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

func (self *Connection) Detail() *transport.ConnectionDetail {
	return self.detail
}

func (self *Connection) PeerCertificates() []*x509.Certificate {
	return self.session.ConnectionState().PeerCertificates
}

// Reader method on the transport.Connection interface.
//
func (self *Connection) Reader() io.Reader {
	return self.stream
}

// Writer method on the transport.Connection interface.
//
func (self *Connection) Writer() io.Writer {
	return self.stream
}

// QUIC doesn't provide an underlying net.Conn
//
func (self *Connection) Conn() net.Conn {
	return nil
}

// SetReadTimeout sets the read timeout to the provided duration.
//
func (self *Connection) SetReadTimeout(t time.Duration) error {
	return self.stream.SetReadDeadline(time.Now().Add(t))
}

// SetWriteTimeout sets the write timeout to the provided duration.
//
func (self *Connection) SetWriteTimeout(t time.Duration) error {
	return self.stream.SetWriteDeadline(time.Now().Add(t))
}

// ClearReadTimeout clears the read time for all current and future reads
//
func (self *Connection) ClearReadTimeout() error {
	var zero time.Time
	return self.stream.SetReadDeadline(zero)
}

// ClearWriteTimeout clears the write timeout for all current and future writes
//
func (self *Connection) ClearWriteTimeout() error {
	var zero time.Time
	return self.stream.SetWriteDeadline(zero)
}

// Close method on the transport.Connection interface.
//
func (self *Connection) Close() error {
	err1 := self.stream.Close()
	err2 := self.session.Close()

	if err1 == nil {
		return err2
	}
	if err2 == nil {
		return err1
	}
	return fmt.Errorf("multiple errors closing quic connection (%s), (%s)", err1, err2)
}
