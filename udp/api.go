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
	"github.com/openziti/foundation/util/mempool"
	"net"
	"time"
)

type WriteQueue interface {
	Accept(mempool.PooledBuffer)
	LocalAddr() net.Addr
}

type NewConnAcceptResult int

const (
	Allow NewConnAcceptResult = iota
	Deny
	AllowDropLRU
)

type NewConnPolicy interface {
	NewConnection(currentCount uint32) NewConnAcceptResult
}

type ConnExpirationPolicy interface {
	IsExpired(now, lastUsed time.Time) bool
	PollFrequency() time.Duration
}

type UnpooledBuffer []byte

func (u UnpooledBuffer) GetPayload() []byte {
	return u
}

func (u UnpooledBuffer) Release() {
	// does nothing
}
