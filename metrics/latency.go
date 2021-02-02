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

package metrics

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/foundation/channel2"
	"github.com/openziti/foundation/util/info"
	"time"
)

const (
	latencyProbeTime = 128
)

// send regular latency probes
//
func ProbeLatency(ch channel2.Channel, histogram Histogram, interval time.Duration) {
	log := pfxlog.ContextLogger(ch.Label())
	log.Debug("started")
	defer log.Debug("exited")
	defer func() {
		histogram.Dispose()
	}()

	lastSend := info.NowInMilliseconds()
	for {
		time.Sleep(interval)
		if ch.IsClosed() {
			return
		}

		now := info.NowInMilliseconds()
		if now-lastSend > 10000 {
			lastSend = now
			request := channel2.NewMessage(channel2.ContentTypeLatencyType, nil)
			request.PutUint64Header(latencyProbeTime, uint64(time.Now().UnixNano()))
			waitCh, err := ch.SendAndWaitWithPriority(request, channel2.High)
			if err != nil {
				log.Errorf("unexpected error sending latency probe (%s)", err)
				continue
			}

			select {
			case response := <-waitCh:
				if response == nil {
					log.Error("wait channel closed")
					return
				}
				if response.ContentType == channel2.ContentTypeResultType {
					result := channel2.UnmarshalResult(response)
					if result.Success {
						if sentTime, ok := response.GetUint64Header(latencyProbeTime); ok {
							latency := time.Now().UnixNano() - int64(sentTime)
							histogram.Update(latency)
						} else {
							log.Error("no send time")
						}
					} else {
						log.Error("failed latency response")
					}
				} else {
					log.Errorf("unexpected latency response [%d]", response.ContentType)
				}
			case <-time.After(time.Second * 5):
				log.Error("latency timeout")
			}
		}
	}
}

type LatencyHandler interface {
	LatencyReported(latency time.Duration)
	ChannelClosed()
}

func AddLatencyProbe(ch channel2.Channel, interval time.Duration, handler LatencyHandler) {
	probe := &latencyProbe{
		handler:  handler,
		ch:       ch,
		interval: interval,
	}
	ch.AddReceiveHandler(probe)
	go probe.run()
}

type latencyProbe struct {
	handler  LatencyHandler
	ch       channel2.Channel
	interval time.Duration
}

func (self *latencyProbe) ContentType() int32 {
	return channel2.ContentTypeLatencyResponseType
}

func (self *latencyProbe) HandleReceive(m *channel2.Message, _ channel2.Channel) {
	if sentTime, ok := m.GetUint64Header(latencyProbeTime); ok {
		latency := time.Duration(time.Now().UnixNano() - int64(sentTime))
		self.handler.LatencyReported(latency)
	} else {
		pfxlog.Logger().Error("no send time on latency response")
	}
}

func (self *latencyProbe) run() {
	log := pfxlog.ContextLogger(self.ch.Label())
	log.Debug("started")
	defer log.Debug("exited")
	defer self.handler.ChannelClosed()

	for !self.ch.IsClosed() {
		request := channel2.NewMessage(channel2.ContentTypeLatencyType, nil)
		request.PutUint64Header(latencyProbeTime, uint64(time.Now().UnixNano()))
		if err := self.ch.SendPrioritizedWithTimeout(request, channel2.High, 10*time.Second); err != nil {
			log.WithError(err).Error("unexpected error sending latency probe")
		}
		time.Sleep(self.interval)
	}
}

func AddLatencyProbeResponder(ch channel2.Channel) {
	responder := &LatencyResponder{
		ch:              ch,
		responseChannel: make(chan *channel2.Message, 1),
	}
	ch.AddReceiveHandler(responder)
	go responder.responseSender()
}

// LatencyResponder responds to latency messages with LatencyResponse messages.
//
type LatencyResponder struct {
	responseChannel chan *channel2.Message
	ch              channel2.Channel
}

func (self *LatencyResponder) ContentType() int32 {
	return channel2.ContentTypeLatencyType
}

func (self *LatencyResponder) HandleReceive(msg *channel2.Message, _ channel2.Channel) {
	if sentTime, found := msg.Headers[latencyProbeTime]; found {
		resp := channel2.NewMessage(channel2.ContentTypeLatencyResponseType, nil)
		resp.Headers[latencyProbeTime] = sentTime
		select {
		case self.responseChannel <- resp:
		default:
		}
	}
}

func (self *LatencyResponder) responseSender() {
	log := pfxlog.ContextLogger(self.ch.Label())
	log.Debug("started")
	defer log.Debug("exited")

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case response := <-self.responseChannel:
			if err := self.ch.SendWithPriority(response, channel2.High); err != nil {
				log.WithError(err).Error("error sending latency response")
				if self.ch.IsClosed() {
					return
				}
			}
		case <-ticker.C:
			if self.ch.IsClosed() {
				return
			}
		}
	}
}
