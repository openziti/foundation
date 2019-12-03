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

package boltz

import (
	"github.com/kataras/go-events"
	"go.etcd.io/bbolt"
)

type MutateContext interface {
	Tx() *bbolt.Tx
	AddEvent(em events.EventEmmiter, name events.EventName, entity BaseEntity)
}

type mutateEvent struct {
	em     events.EventEmmiter
	entity BaseEntity
	name   events.EventName
}

func NewMutateContext(tx *bbolt.Tx) MutateContext {
	context := &mutateContext{tx: tx}
	tx.OnCommit(context.handleCommit)
	return context
}

type mutateContext struct {
	tx     *bbolt.Tx
	events []*mutateEvent
}

func (context *mutateContext) Tx() *bbolt.Tx {
	return context.tx
}

func (context *mutateContext) AddEvent(em events.EventEmmiter, name events.EventName, entity BaseEntity) {
	context.events = append(context.events, &mutateEvent{
		em:     em,
		entity: entity,
		name:   name,
	})
}

func (context *mutateContext) handleCommit() {
	for _, event := range context.events {
		go event.em.Emit(event.name, event.entity)
	}
}
