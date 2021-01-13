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

package transwarptls

import (
	"github.com/openziti/dilithium/cf"
	"github.com/openziti/dilithium/protocol/westlsworld3"
	"github.com/openziti/dilithium/protocol/westworld3"
	"github.com/openziti/foundation/identity/identity"
	"github.com/openziti/foundation/transport"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net"
)

func Dial(endpoint *net.UDPAddr, name string, id *identity.TokenId, tcfg transport.Configuration) (transport.Connection, error) {
	profileId := byte(0)
	if tcfg != nil {
		profile := &westworld3.Profile{}
		if err := profile.Load(cf.MapIToMapS(tcfg)); err != nil {
			return nil, errors.Wrap(err, "load profile")
		}
		newProfileId, err := westworld3.AddProfile(profile)
		if err != nil {
			return nil, errors.Wrap(err, "register profile")
		}
		profileId = newProfileId
	}
	logrus.Infof("westworld3 profile = [\n%s\n]", westworld3.GetProfile(profileId).Dump())

	tlsConfig := id.ClientTLSConfig()
	tlsConfig.ServerName = endpoint.IP.String()
	socket, err := westlsworld3.Dial(endpoint, tlsConfig, profileId)
	if err != nil {
		return nil, errors.Wrap(err, "dial")
	}

	return &Connection{
		detail: &transport.ConnectionDetail{
			Address: "transwarptls:"+endpoint.String(),
			InBound: false,
			Name: name,
		},
		socket: socket,
	}, nil
}