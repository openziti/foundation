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

package udp

import (
	"github.com/netfoundry/ziti-foundation/transport"
	"crypto/x509"
	"io"
	"net"
	"time"
)

const (
	// MaxPacketSize is the maximum size of data that can be in a UDP packet
	MaxPacketSize = 65507
)

type connection struct {
	detail  *transport.ConnectionDetail
	socket  net.Conn
	fullBuf []byte
	copyBuf []byte
}

func (c *connection) Read(p []byte) (int, error) {
	var err error

	if len(c.copyBuf) == 0 {
		var bytesRead int
		bytesRead, err = c.socket.Read(c.fullBuf)
		if bytesRead > 0 {
			c.copyBuf = c.fullBuf[:bytesRead]
		}
	}

	bytesCopied := 0
	if len(c.copyBuf) > 0 {
		bytesCopied = copy(p, c.copyBuf)
		c.copyBuf = c.copyBuf[bytesCopied:]
	}

	return bytesCopied, err
}

func (c *connection) Detail() *transport.ConnectionDetail {
	return c.detail
}

func (c *connection) PeerCertificates() []*x509.Certificate {
	return nil
}

func (c *connection) Reader() io.Reader {
	return c
}

func (c *connection) Writer() io.Writer {
	return c.socket
}

func (c *connection) Conn() net.Conn {
	return c.socket
}

func (c *connection) SetReadTimeout(t time.Duration) error {
	return c.socket.SetReadDeadline(time.Now().Add(t))
}

func (c *connection) SetWriteTimeout(t time.Duration) error {
	return c.socket.SetWriteDeadline(time.Now().Add(t))
}

func (c *connection) Close() error {
	return c.socket.Close()
}
