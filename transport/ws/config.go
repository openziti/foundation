package ws

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type WSConfig struct {
	writeTimeout      time.Duration
	readTimeout       time.Duration
	idleTimeout       time.Duration
	pongTimeout       time.Duration
	pingInterval      time.Duration
	handshakeTimeout  time.Duration
	readBufferSize    int
	writeBufferSize   int
	enableCompression bool
	serverCert        string
	key               string
}

func NewDefaultWSConfig() *WSConfig {
	return &WSConfig{
		writeTimeout:      time.Second * 10,
		readTimeout:       time.Second * 5,
		idleTimeout:       time.Second * 5,
		pongTimeout:       time.Second * 60,
		handshakeTimeout:  time.Second * 10,
		readBufferSize:    4096,
		writeBufferSize:   4096,
		enableCompression: true,
	}
}

func (self *WSConfig) Load(data map[interface{}]interface{}) error {
	if v, found := data["writeTimeout"]; found {
		if i, ok := v.(int); ok {
			self.writeTimeout = time.Second * time.Duration(i)
		} else {
			return errors.New("invalid 'writeTimeout' value")
		}
	}
	if v, found := data["readTimeout"]; found {
		if i, ok := v.(int); ok {
			self.readTimeout = time.Second * time.Duration(i)
		} else {
			return errors.New("invalid 'readTimeout' value")
		}
	}
	if v, found := data["idleTimeout"]; found {
		if i, ok := v.(int); ok {
			self.idleTimeout = time.Second * time.Duration(i)
		} else {
			return errors.New("invalid 'idleTimeout' value")
		}
	}
	if v, found := data["pongTimeout"]; found {
		if i, ok := v.(int); ok {
			self.pongTimeout = time.Second * time.Duration(i)
		} else {
			return errors.New("invalid 'pongTimeout' value")
		}
	}
	if v, found := data["pingInterval"]; found {
		if i, ok := v.(int); ok {
			self.pingInterval = time.Second * time.Duration(i)
		} else {
			return errors.New("invalid 'pingInterval' value")
		}
	} else {
		self.pingInterval = (self.pongTimeout * 9) / 10
	}
	if v, found := data["handshakeTimeout"]; found {
		if i, ok := v.(int); ok {
			self.handshakeTimeout = time.Second * time.Duration(i)
		} else {
			return errors.New("invalid 'handshakeTimeout' value")
		}
	}
	if v, found := data["readBufferSize"]; found {
		if i, ok := v.(int); ok {
			self.readBufferSize = i
		} else {
			return errors.New("invalid 'readBufferSize' value")
		}
	}
	if v, found := data["writeBufferSize"]; found {
		if i, ok := v.(int); ok {
			self.writeBufferSize = i
		} else {
			return errors.New("invalid 'writeBufferSize' value")
		}
	}
	if v, found := data["enableCompression"]; found {
		if i, ok := v.(bool); ok {
			self.enableCompression = i
		} else {
			return errors.New("invalid 'enableCompression' value")
		}
	}
	if v, found := data["server_cert"]; found {
		if s, ok := v.(string); ok {
			self.serverCert = s
		} else {
			return errors.New("invalid 'server_cert' value")
		}
	}
	if v, found := data["key"]; found {
		if s, ok := v.(string); ok {
			self.key = s
		} else {
			return errors.New("invalid 'key' value")
		}
	}

	if len(self.serverCert) == 0 {
		return errors.New("transport.ws.serverCert was not specified'")
	}
	if len(self.key) == 0 {
		return errors.New("transport.ws.key was not specified'")
	}

	return nil
}

func (self *WSConfig) Dump() string {
	out := "ws.Config{\n"
	out += fmt.Sprintf("\t%-30s %d\n", "writeTimeout", self.writeTimeout)
	out += fmt.Sprintf("\t%-30s %d\n", "readTimeout", self.readTimeout)
	out += fmt.Sprintf("\t%-30s %d\n", "idleTimeout", self.idleTimeout)
	out += fmt.Sprintf("\t%-30s %d\n", "pongTimeout", self.pongTimeout)
	out += fmt.Sprintf("\t%-30s %d\n", "handshakeTimeout", self.handshakeTimeout)
	out += fmt.Sprintf("\t%-30s %d\n", "readBufferSize", self.readBufferSize)
	out += fmt.Sprintf("\t%-30s %d\n", "writeBufferSize", self.writeBufferSize)
	out += fmt.Sprintf("\t%-30s %t\n", "enableCompression", self.enableCompression)
	out += fmt.Sprintf("\t%-30s %s\n", "serverCert", self.serverCert)
	out += fmt.Sprintf("\t%-30s %s\n", "key", self.key)
	out += "}"
	return out
}
