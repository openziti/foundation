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

package udp

import (
	"bufio"
	"github.com/openziti/foundation/identity/identity"
	"github.com/openziti/foundation/transport"
	"math"
	"net"
	"time"
)

func Dial(destination *net.UDPAddr, name string, _ *identity.TokenId, timeout time.Duration) (transport.Connection, error) {
	socket, err := net.DialTimeout("udp", destination.String(), timeout)
	if err != nil {
		return nil, err
	}

	return &Connection{
		detail: &transport.ConnectionDetail{
			Address: "udp:" + destination.String(),
			InBound: false,
			Name:    name,
		},
		socket: socket,
		reader: bufio.NewReaderSize(socket, math.MaxUint16),
	}, nil
}
