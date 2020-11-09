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
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/openziti/foundation/transport"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net"
	"sync"
	"sync/atomic"
	"time"
	// _ "unsafe"	// Using go:linkname requires us to import unsafe
)

/**
 *	For the moment, we do not need to exploit the go:linkname mechanism(s) in order to
 *	manipulate portions of the Go runtime, but we leave this code here, commented out,
 *	in case we need to revisit.


// A cipherSuite is a specific combination of key agreement, cipher and MAC function.
type cipherSuite struct {
	id uint16
	// the lengths, in bytes, of the key material needed for each component.
	keyLen int
	macLen int
	ivLen  int
	ka     func(version uint16)
	// flags is a bitmask of the suite* values, above.
	flags  int
	cipher func(key, iv []byte, isRead bool) interface{}
	mac    func(version uint16, macKey []byte)
	aead   func(key, fixedNonce []byte)
}

//go:linkname cipherSuites crypto/tls.cipherSuites
var cipherSuites []*cipherSuite

const (
	// suiteECDHE indicates that the cipher suite involves elliptic curve
	// Diffie-Hellman. This means that it should only be selected when the
	// client indicates that it supports ECC with a curve and point format
	// that we're happy with.
	suiteECDHE = 1 << iota
	// suiteECSign indicates that the cipher suite involves an ECDSA or
	// EdDSA signature and therefore may only be selected when the server's
	// certificate is ECDSA or EdDSA. If this is not set then the cipher suite
	// is RSA based.
	suiteECSign
	// suiteTLS12 indicates that the cipher suite should only be advertised
	// and accepted when using TLS 1.2.
	suiteTLS12
	// suiteSHA384 indicates that the cipher suite uses SHA384 as the
	// handshake hash.
	suiteSHA384
	// suiteDefaultOff indicates that this cipher suite is not included by
	// default.
	suiteDefaultOff
)

*/

// TLS 1.0 - 1.2 cipher suites supported by ziti-sdk-js
const (
	TLS_RSA_WITH_AES_128_CBC_SHA uint16 = 0x002f
	TLS_RSA_WITH_AES_256_CBC_SHA uint16 = 0x0035
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
	detail                   *transport.ConnectionDetail
	cfg                      *WSConfig
	ws                       *websocket.Conn
	tlsConn                  *tls.Conn
	tlsConnHandshakeComplete bool
	log                      *logrus.Entry
	rxbuf                    *safeBuffer
	txbuf                    *safeBuffer
	tlsrxbuf                 *safeBuffer
	tlstxbuf                 *safeBuffer
	done                     chan struct{}
	wmutex                   sync.Mutex
	rmutex                   sync.Mutex
	tlswmutex                sync.Mutex
	tlsrmutex                sync.Mutex
	incoming                 chan transport.Connection
	readCallDepth            int32
	writeCallDepth           int32
}

// Read implements io.Reader by wrapping websocket messages in a buffer.
func (c *Connection) Read(p []byte) (n int, err error) {
	currentDepth := atomic.AddInt32(&c.readCallDepth, 1)
	c.log.Tracef("Read() start currentDepth[%d]", currentDepth)

	if c.rxbuf.Len() == 0 {
		var r io.Reader
		c.rxbuf.Reset()
		if c.tlsConnHandshakeComplete {
			if currentDepth == 1 {
				c.tlsrmutex.Lock()
				defer c.tlsrmutex.Unlock()
			} else if currentDepth == 2 {
				c.rmutex.Lock()
				defer c.rmutex.Unlock()
			}
		} else {
			c.rmutex.Lock()
			defer c.rmutex.Unlock()
		}
		select {
		case <-c.done:
			err = errClosing
		default:
			err = c.ws.SetReadDeadline(time.Now().Add(c.cfg.readTimeout))
			if err == nil {
				if c.tlsConnHandshakeComplete && currentDepth == 1 {
					n, err = c.tlsConn.Read(p)
					atomic.SwapInt32(&c.readCallDepth, (c.readCallDepth - 1))
					c.log.Tracef("Read() end currentDepth[%d]", currentDepth)
					return n, err
				} else {
					_, r, err = c.ws.NextReader()
				}
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

	atomic.SwapInt32(&c.readCallDepth, (c.readCallDepth - 1))

	c.log.Tracef("Read() end currentDepth[%d]", currentDepth)

	return c.rxbuf.Read(p)
}

// Write implements io.Writer and sends binary messages only.
func (c *Connection) Write(p []byte) (n int, err error) {
	return c.write(websocket.BinaryMessage, p)
}

// write wraps the websocket writer.
func (c *Connection) write(messageType int, p []byte) (n int, err error) {
	var txbufLen int
	currentDepth := atomic.AddInt32(&c.writeCallDepth, 1)
	c.log.Debugf("Write() start currentDepth[%d] len[%d]", c.writeCallDepth, len(p))

	if c.tlsConnHandshakeComplete {
		if currentDepth == 1 {
			c.tlswmutex.Lock()
			defer c.tlswmutex.Unlock()
		} else if currentDepth == 2 {
			c.wmutex.Lock()
			defer c.wmutex.Unlock()
		}
	} else {
		c.wmutex.Lock()
		defer c.wmutex.Unlock()
	}

	select {
	case <-c.done:
		err = errClosing
	default:
		var txbufLen int

		if !c.tlsConnHandshakeComplete {
			c.tlstxbuf.Write(p)
			txbufLen = c.tlstxbuf.Len()
			c.log.Tracef("Write() doing TLS handshake (buffering); currentDepth[%d] txbufLen[%d] data[%o]", c.writeCallDepth, txbufLen, p)
		} else if currentDepth == 1 { // if at TLS level (1st level)
			c.tlstxbuf.Write(p)
			txbufLen = c.tlstxbuf.Len()
			c.log.Tracef("Write() doing TLS write; currentDepth[%d] txbufLen[%d] data[%o]", c.writeCallDepth, txbufLen, p)
		} else { // if at websocket level (2nd level)
			c.txbuf.Write(p)
			txbufLen = c.txbuf.Len()
			c.log.Tracef("Write() doing raw write; currentDepth[%d] txbufLen[%d] data[%o]", c.writeCallDepth, txbufLen, p)
		}

		if txbufLen > 20 { // TEMP HACK:  (until I refactor the JS-SDK to accept the message section and data section in separate salvos)
			err = c.ws.SetWriteDeadline(time.Now().Add(c.cfg.writeTimeout))
			if err == nil {
				if !c.tlsConnHandshakeComplete {
					m := make([]byte, txbufLen)
					c.tlstxbuf.Read(m)
					c.log.Tracef("Write() doing TLS handshake (to websocket); currentDepth[%d] txbufLen[%d] data[%o]", c.writeCallDepth, txbufLen, m)
					err = c.ws.WriteMessage(messageType, m)
				} else if currentDepth == 1 {
					m := make([]byte, txbufLen)
					c.tlstxbuf.Read(m)
					c.log.Tracef("Write() doing TLS write (to conn); currentDepth[%d] txbufLen[%d] data[%o]", c.writeCallDepth, txbufLen, m)
					n, err = c.tlsConn.Write(m)
					atomic.SwapInt32(&c.writeCallDepth, (c.writeCallDepth - 1))
					c.log.Tracef("write() end TLS write currentDepth[%d]", c.writeCallDepth)
					return n, err
				} else {
					m := make([]byte, txbufLen)
					c.txbuf.Read(m)
					c.log.Debugf("Write() doing raw write (to websocket); currentDepth[%d] len[%d]", c.writeCallDepth, len(m))
					err = c.ws.WriteMessage(messageType, m)
				}
			}
		}
	}
	if err == nil {
		n = txbufLen
	}
	atomic.SwapInt32(&c.writeCallDepth, (c.writeCallDepth - 1))
	c.log.Tracef("Write() end currentDepth[%d]", c.writeCallDepth)

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
			c.log.Trace("sending websocket Ping")
			if _, err := c.write(websocket.PingMessage, []byte{}); err != nil {
				_ = c.Close()
			}
		}
	}
}

/**
 *	See above note re go:linkname
 *
func (c *Connection) patchCipherSuites() {
	c.log.Debug("patchCipherSuites dump: v----------------------------------------------------------")
	for _, cipherSuite := range cipherSuites {
		if cipherSuite.id == TLS_RSA_WITH_AES_128_CBC_SHA {
			c.log.Debug("cipherSuite: TLS_RSA_WITH_AES_128_CBC_SHA before: ", cipherSuite)
			cipherSuite.flags = suiteTLS12 | suiteECDHE
			c.log.Debug("cipherSuite: TLS_RSA_WITH_AES_128_CBC_SHA after: ", cipherSuite)
		}
		if cipherSuite.id == TLS_RSA_WITH_AES_256_CBC_SHA {
			c.log.Debug("cipherSuite: TLS_RSA_WITH_AES_256_CBC_SHA before: ", cipherSuite)
			cipherSuite.flags = suiteTLS12 | suiteECDHE
			c.log.Debug("cipherSuite: TLS_RSA_WITH_AES_256_CBC_SHA after: ", cipherSuite)
		}
	}
	c.log.Debug("patchCipherSuites dump: ^----------------------------------------------------------")
}
*/

// tlsHandshake wraps the websocket in a TLS server.
func (c *Connection) tlsHandshake() error {
	var err error
	var serverCertPEM []byte
	var keyPEM []byte

	//patchCipherSuites()

	if serverCertPEM, err = ioutil.ReadFile(c.cfg.serverCert); err != nil {
		c.log.Error(err)
		_ = c.Close()
		return err
	}

	if keyPEM, err = ioutil.ReadFile(c.cfg.key); err != nil {
		c.log.Error(err)
		_ = c.Close()
		return err
	}

	cert, err := tls.X509KeyPair(serverCertPEM, keyPEM)
	if err != nil {
		c.log.Error(err)
		_ = c.Close()
		return err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(serverCertPEM)

	cfg := &tls.Config{
		ClientCAs:    caCertPool,
		Certificates: []tls.Certificate{cert},
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
		ClientAuth:                  tls.RequireAndVerifyClientCert,
		MinVersion:                  tls.VersionTLS11,
		PreferServerCipherSuites:    true,
		DynamicRecordSizingDisabled: true,
	}

	c.tlsConn = tls.Server(c, cfg)
	if err = c.tlsConn.Handshake(); err != nil {
		if err != nil {
			c.log.Error(err)
			_ = c.Close()
			return err
		}
	}

	c.tlsConnHandshakeComplete = true

	c.log.Debug("TLS Handshake completed successfully")

	return nil
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
	if self.tlsConnHandshakeComplete {
		return self.tlsConn.ConnectionState().PeerCertificates
	} else {
		return nil
	}
}

func (self *Connection) Reader() io.Reader {
	return self
}

func (self *Connection) Writer() io.Writer {
	return self
}

func (self *Connection) Conn() net.Conn {
	self.log.Debug("Conn() entered, returning TLS connection that wraps the websocket")
	return self.tlsConn // Obtain the TLS connection that wraps the websocket
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

func (self *Connection) LocalAddr() net.Addr {
	return self.ws.UnderlyingConn().LocalAddr()
}
func (self *Connection) RemoteAddr() net.Addr {
	return self.ws.UnderlyingConn().RemoteAddr()
}
func (self *Connection) SetDeadline(t time.Time) error {
	return self.ws.UnderlyingConn().SetDeadline(t)
}
func (self *Connection) SetReadDeadline(t time.Time) error {
	return self.ws.UnderlyingConn().SetReadDeadline(t)
}
func (self *Connection) SetWriteDeadline(t time.Time) error {
	return self.ws.UnderlyingConn().SetWriteDeadline(t)
}
