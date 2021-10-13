package udp

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/foundation/util/concurrenz"
	"github.com/openziti/foundation/util/info"
	"github.com/openziti/foundation/util/mempool"
	"github.com/pkg/errors"
	"io"
	"net"
	"time"
)

func Listen(network string, addr *net.UDPAddr) (net.Listener, error) {
	return ListenWithPolicies(network, addr, unlimitedConnections{}, defaultExpirationPolicy{})
}

func ListenWithPolicies(network string, addr *net.UDPAddr, newConnPolicy NewConnPolicy, expirationPolicy ConnExpirationPolicy) (net.Listener, error) {

	socket, err := net.ListenUDP(network, addr)
	if err != nil {
		return nil, err
	}
	listener := &udpListener{
		addr:             addr,
		socket:           socket,
		acceptChannel:    make(chan net.Conn, 1),
		eventC:           make(chan listenerEvent, 16),
		connMap:          map[string]*udpConn{},
		newConnPolicy:    newConnPolicy,
		expirationPolicy: expirationPolicy,
	}
	go listener.readLoop()
	go listener.eventLoop()

	return listener, nil
}

type udpListener struct {
	addr             *net.UDPAddr
	socket           *net.UDPConn
	closed           concurrenz.AtomicBoolean
	acceptChannel    chan net.Conn
	eventC           chan listenerEvent
	connMap          map[string]*udpConn
	newConnPolicy    NewConnPolicy
	expirationPolicy ConnExpirationPolicy
}

func (self *udpListener) Accept() (net.Conn, error) {
	select {
	case conn, ok := <-self.acceptChannel:
		if !ok {
			return nil, errors.New("listener closed")
		}
		return conn, nil
	}
}

func (self *udpListener) Close() error {
	if self.closed.CompareAndSwap(false, true) {
		return self.socket.Close()
	}
	return nil
}

func (self *udpListener) Addr() net.Addr {
	return self.addr
}

func (self *udpListener) readLoop() {
	log := pfxlog.Logger()
	log.Info("starting udp listener read loop")
	defer log.Info("stopping udp listener read loop")

	bufPool := mempool.NewPool(16, info.MaxUdpPacketSize)
	for !self.closed.Get() {
		buf := bufPool.AcquireBuffer()
		n, srcAddr, err := self.socket.ReadFromUDP(buf.Buf)
		if err != nil {
			log.WithError(err).Error("failure while reading udp message. stopping UDP read loop")
			self.eventC <- errorEvent{error: err}
			return
		}

		log.Debugf("read %v bytes from udp, queuing", len(buf.GetPayload()))
		buf.Buf = buf.Buf[:n]
		self.eventC <- &udpReadEvent{
			buf:     buf,
			srcAddr: srcAddr,
		}
	}
}

func (self *udpListener) eventLoop() {
	log := pfxlog.Logger()
	log.Info("starting udp listener event loop")
	defer log.Info("shutting down udp listener event loop")

	timer := time.NewTicker(self.expirationPolicy.PollFrequency())
	defer timer.Stop()

	for {
		select {
		case event, ok := <-self.eventC:
			if !ok {
				return
			}

			err := event.handle(self)
			if err == io.EOF {
				log.Errorf("EOF detected. stopping UDP event loop")
				return
			}
			if err != nil {
				log.Errorf("error while handling udp event: %v", err)
			}
		case <-timer.C:
			self.dropExpired()
		}
	}
}

func (self *udpListener) getWriteQueue(srcAddr net.Addr) WriteQueue {
	pfxlog.Logger().Debugf("Looking up address %v", srcAddr)
	result := self.connMap[srcAddr.String()]
	if result == nil {
		return nil
	}
	return result
}

func (self *udpListener) createWriteQueue(srcAddr net.Addr) (WriteQueue, error) {
	switch self.newConnPolicy.NewConnection(uint32(len(self.connMap))) {
	case AllowDropLRU:
		self.dropLRU()
	case Deny:
		return nil, errors.New("max connections exceeded")
	}
	conn := &udpConn{
		readC:       make(chan mempool.PooledBuffer),
		closeNotify: make(chan struct{}),
		srcAddr:     srcAddr,
		writeConn:   self.socket,
	}
	conn.markUsed()
	self.connMap[srcAddr.String()] = conn

	self.acceptChannel <- conn

	pfxlog.Logger().WithField("udpConnId", srcAddr.String()).Debug("created new virtual UDP connection")

	return conn, nil
}

func (self *udpListener) dropExpired() {
	log := pfxlog.Logger()
	now := time.Now()
	for key, conn := range self.connMap {
		if conn.closed.Get() {
			delete(self.connMap, conn.srcAddr.String())
		}
		if self.expirationPolicy.IsExpired(now, conn.GetLastUsed()) {
			log.WithField("udpConnId", key).Debug("connection expired. removing from UDP vconn manager")
			self.close(conn)
		}
	}
}

func (self *udpListener) dropLRU() {
	if len(self.connMap) < 1 {
		return
	}
	var oldest *udpConn
	for _, value := range self.connMap {
		if oldest == nil {
			oldest = value
		} else if oldest.GetLastUsed().After(value.GetLastUsed()) {
			oldest = value
		}
	}
	self.close(oldest)
}

func (self *udpListener) close(conn *udpConn) {
	_ = conn.Close()
	delete(self.connMap, conn.srcAddr.String())
}

type udpReadEvent struct {
	buf     *mempool.DefaultPooledBuffer
	srcAddr net.Addr
}

func (event *udpReadEvent) handle(listener *udpListener) error {
	log := pfxlog.Logger()

	writeQueue := listener.getWriteQueue(event.srcAddr)

	if writeQueue == nil {
		log.Debugf("received connection for %v --> %v", event.srcAddr, listener.socket.LocalAddr())
		var err error
		writeQueue, err = listener.createWriteQueue(event.srcAddr)
		if err != nil {
			return err
		}
	}

	log.Tracef("received %v bytes from %v", len(event.buf.Buf), writeQueue.LocalAddr())
	writeQueue.Accept(event.buf)

	return nil
}

type listenerEvent interface {
	handle(listener *udpListener) error
}

type errorEvent struct {
	error
}

func (self errorEvent) handle(*udpListener) error {
	return self.error
}
