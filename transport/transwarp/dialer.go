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

package transwarp

import (
	"github.com/michaelquigley/dilithium/protocol/westworld2"
	"github.com/openziti/foundation/transport"
	"github.com/pkg/errors"
	"net"
)

func Dial(endpoint *net.UDPAddr, name string) (transport.Connection, error) {
	socket, err := westworld2.Dial(endpoint, westworld2.NewDefaultConfig())
	if err != nil {
		return nil, errors.Wrap(err, "dial")
	}

	return &Connection{
		detail: &transport.ConnectionDetail{
			Address: "transwarp:" + endpoint.String(),
			InBound: false,
			Name:    name,
		},
		socket: socket,
	}, nil
}
