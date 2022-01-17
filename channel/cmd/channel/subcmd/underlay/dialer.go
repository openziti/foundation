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

package underlay

import (
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/foundation/channel2"
	"github.com/openziti/foundation/identity/dotziti"
	"github.com/openziti/foundation/transport"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	dialer.Flags().StringVarP(&dialerIdentity, "identity", "i", "default", ".ziti Identity")
	dialer.Flags().StringVarP(&dialerEndpointAddress, "endpoint", "e", "tcp:127.0.0.1:9999", "Endpoint address")
	dialer.Flags().IntVarP(&dialerMessageCount, "count", "c", 100, "Message count")
	dialer.Flags().IntVarP(&dialerSleepMs, "sleep", "s", 50, "Sleep (ms) between messages")
	dialer.Flags().StringVarP(&dialerUnderlay, "underlay", "u", "classic", "use underlay <classic|reconnecting>")
	underlay.AddCommand(dialer)
}

var dialer = &cobra.Command{
	Use:   "dialer",
	Short: "Launch a dialer",
	Run:   runDialer,
}
var dialerIdentity string
var dialerEndpointAddress string
var dialerMessageCount int
var dialerSleepMs int
var dialerUnderlay string

func runDialer(_ *cobra.Command, _ []string) {
	_, id, err := dotziti.LoadIdentity(dialerIdentity)
	if err != nil {
		panic(err)
	}

	endpoint, err := transport.ParseAddress(dialerEndpointAddress)
	if err != nil {
		panic(err)
	}

	log := pfxlog.Logger()
	options := channel2.DefaultOptions()
	options.BindHandlers = []channel2.BindHandler{&bindHandler{}}

	var dialer channel2.UnderlayFactory
	switch dialerUnderlay {
	case "classic":
		dialer = channel2.NewClassicDialer(id, endpoint, nil)
	case "reconnecting":
		dialer = channel2.NewReconnectingDialer(id, endpoint, nil)
	default:
		panic(fmt.Errorf("unknown underlay [%s]", dialerUnderlay))
	}

	ch, err := channel2.NewChannel("channel", dialer, options)
	if err != nil {
		panic(err)
	}
	log.Infof("channel label = [%s]", ch.Label())

	for i := 0; i < dialerMessageCount; i++ {
		if err := ch.Send(newMessage(i)); err != nil {
			panic(err)
		}
		log.Infof("send = [%d]", i)
		time.Sleep(time.Duration(dialerSleepMs) * time.Millisecond)
	}

	if err := ch.Close(); err != nil {
		panic(err)
	}
}
