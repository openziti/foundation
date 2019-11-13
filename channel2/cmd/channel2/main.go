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

package main

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/netfoundry/ziti-foundation/channel2/cmd/channel2/subcmd"
	_ "github.com/netfoundry/ziti-foundation/channel2/cmd/channel2/subcmd/memory"
	_ "github.com/netfoundry/ziti-foundation/channel2/cmd/channel2/subcmd/underlay"
	"github.com/netfoundry/ziti-foundation/transport"
	"github.com/netfoundry/ziti-foundation/transport/quic"
	"github.com/netfoundry/ziti-foundation/transport/tcp"
	"github.com/netfoundry/ziti-foundation/transport/tls"
	"github.com/sirupsen/logrus"
)

func init() {
	transport.AddAddressParser(quic.AddressParser{})
	transport.AddAddressParser(tcp.AddressParser{})
	transport.AddAddressParser(tls.AddressParser{})
	pfxlog.Global(logrus.InfoLevel)
	pfxlog.SetPrefix("bitbucket.org/netfoundry/")
}

func main() {
	if err := subcmd.Root.Execute(); err != nil {
		panic(err)
	}
}
