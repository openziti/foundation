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
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/dilithium/cf"
	"github.com/openziti/dilithium/protocol/westlsworld3"
	"github.com/openziti/dilithium/protocol/westworld3"
	"github.com/openziti/foundation/identity/identity"
	"github.com/openziti/foundation/transport"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"net"
)

func Listen(bind *net.UDPAddr, name string, id *identity.TokenId, incoming chan transport.Connection, tcfg transport.Configuration) (io.Closer, error) {
	log := pfxlog.ContextLogger(name + "/transwarptls:" + bind.String())

	profileId := byte(0)
	if tcfg != nil {
		profile := westworld3.NewBaselineProfile()
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

	listener, err := westlsworld3.Listen(bind, id.ServerTLSConfig(), profileId)
	if err != nil {
		return nil, err
	}

	go acceptLoop(log, name, listener, incoming)

	return listener, nil
}

func acceptLoop(log *logrus.Entry, name string, listener net.Listener, incoming chan transport.Connection) {
	defer log.Error("exited")

	for {
		socket, err := listener.Accept()
		if err != nil {
			if netErr, ok := err.(net.Error); ok && !netErr.Temporary() {
				log.WithField("err", err).Error("accept failed. failure not recoverable. exiting listen loop")
				return
			}
			log.WithField("err", err).Error("accept failed")

		} else {
			connection := &Connection{
				detail: &transport.ConnectionDetail{
					Address: "transwarptls:" + socket.RemoteAddr().String(),
					InBound: true,
					Name:    name,
				},
				socket: socket,
			}
			incoming <- connection

			log.WithField("addr", socket.RemoteAddr().String()).Info("accepted connection")
		}
	}
}