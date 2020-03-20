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
	"fmt"
	"github.com/netfoundry/ziti-foundation/transport"
	"net"
)

// Dial attempts to dial a UDP endpoint and create a connection to it
func Dial(destination, name string) (transport.Connection, error) {
	localAddress, err := net.ResolveUDPAddr("udp", ":0")
	if err != nil {
		return nil, fmt.Errorf("error resolving local address (%w)", err)
	}

	destinationAddress, err := net.ResolveUDPAddr("udp", destination)
	if err != nil {
		return nil, fmt.Errorf("error resolving destination address (%w)", err)
	}

	socket, err := net.DialUDP("udp", localAddress, destinationAddress)
	if err != nil {
		return nil, err
	}

	return &connection{
		detail: &transport.ConnectionDetail{
			Address: "udp:" + destination,
			InBound: false,
			Name:    name,
		},
		socket:  socket,
		fullBuf: make([]byte, MaxPacketSize),
	}, nil
}
