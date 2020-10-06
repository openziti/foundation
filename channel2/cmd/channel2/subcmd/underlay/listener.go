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
	listener.Flags().StringVarP(&listenerIdentity, "identity", "i", "default", ".ziti Idenitity")
	listener.Flags().StringVarP(&listenerEndpointAddress, "endpoint", "e", "tcp:127.0.0.1:9999", "Endpoint address")
	listener.Flags().IntVarP(&listenerMessageCount, "count", "c", 200, "Message count")
	listener.Flags().IntVarP(&listenerSleepMs, "sleep", "s", 50, "Sleep (ms) between messages")
	listener.Flags().StringVarP(&listenerUnderlay, "underlay", "u", "classic", "use underlauy <classic|reconnecting>")
	underlay.AddCommand(listener)
}

var listener = &cobra.Command{
	Use:   "listener",
	Short: "Launch a listener",
	Run:   runListener,
}
var listenerIdentity string
var listenerEndpointAddress string
var listenerMessageCount int
var listenerSleepMs int
var listenerUnderlay string

func runListener(_ *cobra.Command, _ []string) {
	_, id, err := dotziti.LoadIdentity(listenerIdentity)
	if err != nil {
		panic(err)
	}

	endpoint, err := transport.ParseAddress(listenerEndpointAddress)
	if err != nil {
		panic(err)
	}

	var listener channel2.UnderlayListener
	switch listenerUnderlay {
	case "classic":
		listener = channel2.NewClassicListener(id, endpoint, channel2.DefaultConnectOptions(), nil)
	case "reconnecting":
		panic("not implemented")
	default:
		panic(fmt.Errorf("unknown underlay [%s]", listenerUnderlay))
	}

	if err := listener.Listen(); err != nil {
		panic(err)
	}

	log := pfxlog.Logger()
	options := channel2.DefaultOptions()
	options.BindHandlers = []channel2.BindHandler{&bindHandler{}}

	for {
		ch, err := channel2.NewChannel("channel2", listener, options)
		if err != nil {
			panic(err)
		}
		log.Infof("channel label = [%s]", ch.Label())

		go handleChannel(ch)
	}
}

func handleChannel(ch channel2.Channel) {
	log := pfxlog.ContextLogger(ch.Label())
	for i := 0; i < listenerMessageCount; i++ {
		if err := ch.Send(newMessage(i)); err != nil {
			log.Errorf("error sending (%s)", err)
			break
		}
		log.Infof("send = [%d]", i)
		time.Sleep(time.Duration(listenerSleepMs) * time.Millisecond)
	}
}
