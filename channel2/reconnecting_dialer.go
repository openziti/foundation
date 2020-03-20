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

package channel2

import (
	"github.com/netfoundry/ziti-foundation/identity/identity"
	"github.com/netfoundry/ziti-foundation/transport"
	"errors"
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"sync"
	"time"
)

type reconnectingDialer struct {
	identity      *identity.TokenId
	endpoint      transport.Address
	headers       map[int32][]byte
	reconnectLock sync.Mutex
}

func NewReconnectingDialer(identity *identity.TokenId, endpoint transport.Address, headers map[int32][]byte) UnderlayFactory {
	return &reconnectingDialer{
		identity: identity,
		endpoint: endpoint,
		headers:  headers,
	}
}

func (dialer *reconnectingDialer) Create() (Underlay, error) {
	log := pfxlog.ContextLogger(dialer.endpoint.String())
	log.Debug("started")
	defer log.Debug("exited")

	retryCount := 0
	version := uint32(2)

	for {
		peer, err := dialer.endpoint.Dial("reconnecting", dialer.identity)
		if err != nil {
			return nil, err
		}

		impl := newReconnectingImpl(peer, dialer)
		impl.setProtocolVersion(version)

		if err := dialer.sendHello(impl); err != nil {
			if retryCount == 0 {
				return nil, err
			}
			retryCount++
			version, _ = getRetryVersion(err)
			log.Warnf("Retrying initial dial with protocol version %v", version)
			continue
		}
		return impl, nil
	}
}

func (dialer *reconnectingDialer) Reconnect(impl *reconnectingImpl) error {
	log := pfxlog.ContextLogger(impl.Label() + " @" + dialer.endpoint.String())
	log.Debug("starting")
	defer log.Debug("exiting")

	dialer.reconnectLock.Lock()
	defer dialer.reconnectLock.Unlock()

	if err := impl.pingInstance(); err != nil {
		log.Errorf("unable to ping (%s)", err)
		for i := 0; true; i++ {
			peer, err := dialer.endpoint.Dial("reconnecting", dialer.identity)
			if err == nil {
				impl.peer = peer
				if err := dialer.sendHello(impl); err == nil {
					return nil
				} else {
					if version, ok := getRetryVersion(err); ok {
						impl.setProtocolVersion(version)
					}
					log.Errorf("hello attempt [#%d] failed (%s)", i+1, err)
					time.Sleep(5 * time.Second)
				}

			} else {
				log.Errorf("reconnection attempt [#%d] failed (%s)", i+1, err)
				time.Sleep(5 * time.Second)
			}
		}
	}
	return nil
}

func (dialer *reconnectingDialer) sendHello(impl *reconnectingImpl) error {
	log := pfxlog.ContextLogger(impl.Label())
	defer log.Debug("exited")
	log.Debug("started")

	request := NewHello(dialer.identity.Token, dialer.headers)
	request.sequence = HelloSequence
	if impl.connectionId != "" {
		request.Headers[ConnectionIdHeader] = []byte(impl.connectionId)
		log.Debugf("adding connectionId header [%s]", impl.connectionId)
	}
	if err := impl.Tx(request); err != nil {
		_ = impl.peer.Close()
		return err
	}

	response, err := impl.Rx()
	if err != nil {
		return err
	}
	if !response.IsReplyingTo(request.sequence) || response.ContentType != ContentTypeResultType {
		return fmt.Errorf("channel synchronization error, expected %v, got %v", request.sequence, response.ReplyFor())
	}
	result := UnmarshalResult(response)
	if !result.Success {
		return errors.New(result.Message)
	}
	impl.connectionId = string(response.Headers[ConnectionIdHeader])

	return nil
}
