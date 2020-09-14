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

package wss

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/openziti/foundation/transport"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"sync"
	"time"
)

var (
	errClosing = errors.New(`Closing`)
)

// safeBuffer adds thread-safety to *bytes.Buffer
type safeBuffer struct {
	buf *bytes.Buffer
	log *logrus.Entry
	sync.Mutex
}

// Read reads the next len(p) bytes from the buffer or until the buffer is drained.
func (s *safeBuffer) Read(p []byte) (int, error) {
	s.Lock()
	defer s.Unlock()
	return s.buf.Read(p)
}

// Write appends the contents of p to the buffer.
func (s *safeBuffer) Write(p []byte) (int, error) {
	s.Lock()
	defer s.Unlock()
	return s.buf.Write(p)
}

// Len returns the number of bytes of the unread portion of the buffer.
func (s *safeBuffer) Len() int {
	s.Lock()
	defer s.Unlock()
	return s.buf.Len()
}

// Reset resets the buffer to be empty.
func (s *safeBuffer) Reset() {
	s.Lock()
	s.buf.Reset()
	s.Unlock()
}

// Connection wraps gorilla websocket to provide io.ReadWriteCloser
type Connection struct {
	detail *transport.ConnectionDetail
	cfg    *WSSConfig
	ws     *websocket.Conn
	log    *logrus.Entry
	rxbuf  *safeBuffer
	txbuf  *safeBuffer
	done   chan struct{}
	wmutex sync.Mutex
	rmutex sync.Mutex
}

// Read implements io.Reader by wrapping websocket messages in a buffer.
func (c *Connection) Read(p []byte) (n int, err error) {
	if c.rxbuf.Len() == 0 {
		var r io.Reader
		c.rxbuf.Reset()
		c.rmutex.Lock()
		defer c.rmutex.Unlock()
		select {
		case <-c.done:
			err = errClosing
		default:
			err = c.ws.SetReadDeadline(time.Now().Add(c.cfg.pongTimeout))
			if err == nil {
				_, r, err = c.ws.NextReader()
			}
		}
		if err != nil {
			return n, err
		}
		_, err = io.Copy(c.rxbuf, r)
		if err != nil {
			return n, err
		}
	}

	return c.rxbuf.Read(p)
}

// Write implements io.Writer and sends binary messages only.
func (c *Connection) Write(p []byte) (n int, err error) {
	return c.write(websocket.BinaryMessage, p)
}

// write wraps the websocket writer.
func (c *Connection) write(messageType int, p []byte) (n int, err error) {
	var txbufLen int
	c.wmutex.Lock()
	defer c.wmutex.Unlock()
	select {
	case <-c.done:
		err = errClosing
	default:
		c.txbuf.Write(p)
		txbufLen = c.txbuf.Len()
		if txbufLen > 20 { // TEMP HACK:  (until I refactor the JS-SDK to accept the message section and data section in separate salvos)
			err = c.ws.SetWriteDeadline(time.Now().Add(c.cfg.writeTimeout))
			if err == nil {
				m := make([]byte, txbufLen)
				c.txbuf.Read(m)
				err = c.ws.WriteMessage(messageType, m)
			}
		}
	}
	if err == nil {
		n = txbufLen
	}
	return n, err
}

// Close implements io.Closer and closes the underlying connection.
func (c *Connection) Close() error {
	c.rmutex.Lock()
	c.wmutex.Lock()
	defer func() {
		c.rmutex.Unlock()
		c.wmutex.Unlock()
	}()
	select {
	case <-c.done:
		return errClosing
	default:
		close(c.done)
	}
	return c.ws.Close()
}

// pinger sends ping messages on an interval for client keep-alive.
func (c *Connection) pinger() {
	ticker := time.NewTicker(c.cfg.pingInterval)
	defer ticker.Stop()
	for {
		select {
		case <-c.done:
			return
		case <-ticker.C:
			if _, err := c.write(websocket.PingMessage, []byte{}); err != nil {
				_ = c.Close()
			}
		}
	}
}

// newSafeBuffer instantiates a new safeBuffer
func newSafeBuffer(log *logrus.Entry) *safeBuffer {
	return &safeBuffer{
		buf: bytes.NewBuffer(nil),
		log: log,
	}
}

func (self *Connection) Detail() *transport.ConnectionDetail {
	return self.detail
}

func (self *Connection) PeerCertificates() []*x509.Certificate {
	var tlsConn (*tls.Conn) = self.ws.UnderlyingConn().(*tls.Conn)
	return tlsConn.ConnectionState().PeerCertificates
}

func (self *Connection) Reader() io.Reader {
	return self
}

func (self *Connection) Writer() io.Writer {
	return self
}

func (self *Connection) Conn() net.Conn {
	return self.ws.UnderlyingConn() // Obtain the socket underneath the websocket

}

func (self *Connection) SetReadTimeout(t time.Duration) error {
	return self.ws.UnderlyingConn().SetReadDeadline(time.Now().Add(t))
}

func (self *Connection) SetWriteTimeout(t time.Duration) error {
	return self.ws.UnderlyingConn().SetWriteDeadline(time.Now().Add(t))
}

// ClearReadTimeout clears the read time for all current and future reads
//
func (self *Connection) ClearReadTimeout() error {
	var zero time.Time
	return self.ws.UnderlyingConn().SetReadDeadline(zero)
}

// ClearWriteTimeout clears the write timeout for all current and future writes
//
func (self *Connection) ClearWriteTimeout() error {
	var zero time.Time
	return self.ws.UnderlyingConn().SetWriteDeadline(zero)
}
