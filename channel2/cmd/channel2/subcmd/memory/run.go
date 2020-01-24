/*
	Copyright 2019 NetFoundry, Inc.

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
	"github.com/netfoundry/ziti-foundation/channel2"
	"github.com/netfoundry/ziti-foundation/channel2/cmd/channel2/subcmd"
	"github.com/netfoundry/ziti-foundation/identity/dotziti"
	"github.com/netfoundry/ziti-foundation/identity/identity"
	"github.com/michaelquigley/pfxlog"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	memory.Flags().StringVarP(&memoryIdentity, "identity", "i", "default", ".ziti Identity")
	memory.Flags().IntVarP(&memoryCount, "count", "c", 100, "number of messages to send")
	subcmd.Root.AddCommand(memory)
}

var memory = &cobra.Command{
	Use:   "memory",
	Short: "colocated dialer/listener using memory underlay",
	Run:   runMemory,
}
var memoryIdentity string
var memoryCount int

var listenerDone = make(chan struct{})
var dialerDone = make(chan struct{})

func runMemory(_ *cobra.Command, _ []string) {
	_, id, err := dotziti.LoadIdentity(memoryIdentity)
	if err != nil {
		panic(err)
	}

	ctx := channel2.NewMemoryContext()
	go handleDialer(id, ctx)
	go runListener(id, ctx)

	<-listenerDone
	<-dialerDone

	/*
	 * Time for the rxer/txer goroutines to exit.
	 */
	time.Sleep(1 * time.Second)
	/* */
}

func handleDialer(identity *identity.TokenId, ctx *channel2.MemoryContext) {
	log := pfxlog.Logger()

	options := channel2.DefaultOptions()
	options.BindHandlers = []channel2.BindHandler{&bindHandler{}}

	dialer := channel2.NewMemoryDialer(identity, nil, ctx)
	ch, err := channel2.NewChannel("memory", dialer, options)
	if err != nil {
		panic(err)
	}
	log.Infof("channel label = [%s]", ch.Label())

	log = pfxlog.ContextLogger(ch.Label())
	for i := 0; i < memoryCount; i++ {
		if err := ch.Send(newMessage(i)); err != nil {
			log.Errorf("error sending (%s)", err)
			break
		}
		log.Infof("send = [%d]", i)
	}

	if err := ch.Close(); err != nil {
		panic(err)
	}

	close(dialerDone)
}

func runListener(identity *identity.TokenId, ctx *channel2.MemoryContext) {
	log := pfxlog.Logger()

	listener := channel2.NewMemoryListener(identity, ctx)
	if err := listener.Listen(); err != nil {
		panic(err)
	}

	options := channel2.DefaultOptions()
	options.BindHandlers = []channel2.BindHandler{&bindHandler{}}

	ch, err := channel2.NewChannel("memory", listener, options)
	if err != nil {
		panic(err)
	}
	log.Infof("channel label = [%s]", ch.Label())

	go handleListener(ch)

	if err := listener.Close(); err != nil {
		log.Errorf("error closing listener (%s)", err)
	}
}

func handleListener(ch channel2.Channel) {
	log := pfxlog.ContextLogger(ch.Label())
	for i := 0; i < memoryCount; i++ {
		if err := ch.Send(newMessage(i)); err != nil {
			log.Errorf("error sending (%s)", err)
			break
		}
		log.Infof("send = [%d]", i)
	}

	if err := ch.Close(); err != nil {
		panic(err)
	}

	close(listenerDone)
}