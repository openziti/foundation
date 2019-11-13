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

package udp

import (
	"github.com/netfoundry/ziti-foundation/identity/identity"
	"github.com/netfoundry/ziti-foundation/transport"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var _ transport.Address = (*address)(nil) // enforce that address implements transport.Address

type address struct {
	hostname string
	port     uint16
}

func (a address) Dial(name string, i *identity.TokenId) (transport.Connection, error) {
	return Dial(a.bindableAddress(), name)
}

func (a address) Listen(name string, i *identity.TokenId, incoming chan transport.Connection) (io.Closer, error) {
	panic(errors.New("UDP listen not supported"))
}

func (a address) MustListen(name string, i *identity.TokenId, incoming chan transport.Connection) io.Closer {
	closer, err := a.Listen(name, i, incoming)
	if err != nil {
		panic(err)
	}
	return closer
}

func (a address) String() string {
	return fmt.Sprintf("udp:%s", a.bindableAddress())
}

func (a address) bindableAddress() string {
	return fmt.Sprintf("%s:%d", a.hostname, a.port)
}

// AddressParser is a parser for UDP addresses
type AddressParser struct{}

// Parse attempts to parse a UDP address from a string
func (ap AddressParser) Parse(s string) (transport.Address, error) {
	tokens := strings.Split(s, ":")
	if len(tokens) < 2 {
		return nil, errors.New("invalid format")
	}

	if tokens[0] == "udp" {
		if len(tokens) != 3 {
			return nil, errors.New("invalid format")
		}

		port, err := strconv.ParseUint(tokens[2], 10, 16)
		if err != nil {
			return nil, err
		}

		return &address{hostname: tokens[1], port: uint16(port)}, nil
	}

	return nil, errors.New("invalid format")
}
