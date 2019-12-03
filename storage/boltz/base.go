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
	"github.com/netfoundry/ziti-foundation/storage/ast"
	"github.com/kataras/go-events"
	"go.etcd.io/bbolt"
	"io"
	"time"
)

const (
	RootBucket    = "ziti"
	IndexesBucket = "indexes"
)

const (
	EventCreate events.EventName = "CREATE"
	EventDelete events.EventName = "DELETE"
	EventUpdate events.EventName = "UPDATE"
)

type Db interface {
	io.Closer
	Update(fn func(tx *bbolt.Tx) error) error
	View(fn func(tx *bbolt.Tx) error) error
	RootBucket(tx *bbolt.Tx) (*bbolt.Bucket, error)
}

type ListStore interface {
	ast.SymbolTypes

	GetEntityType() string
	GetEntitiesBucket(tx *bbolt.Tx) *TypedBucket
	GetOrCreateEntitiesBucket(tx *bbolt.Tx) *TypedBucket
	GetEntityBucket(tx *bbolt.Tx, id []byte) *TypedBucket
	GetOrCreateEntityBucket(tx *bbolt.Tx, id []byte) *TypedBucket
	GetValue(tx *bbolt.Tx, id []byte, path ...string) []byte
	GetValueCursor(tx *bbolt.Tx, id []byte, path ...string) *bbolt.Cursor
	IsChildStore() bool
	IsEntityPresent(tx *bbolt.Tx, id string) bool

	GetSymbol(name string) EntitySymbol
	GrantSymbols(child ListStore)
	inheritSymbol(symbol EntitySymbol)
	AddIdSymbol(name string, nodeType ast.NodeType) EntitySymbol
	AddSymbol(name string, nodeType ast.NodeType, path ...string) EntitySymbol
	AddFkSymbol(name string, linkedType ListStore, path ...string) EntitySymbol
	AddSymbolWithKey(name string, nodeType ast.NodeType, key string, path ...string) EntitySymbol
	AddFkSymbolWithKey(name string, key string, linkedType ListStore, path ...string) EntitySymbol
	AddMapSymbol(name string, nodeType ast.NodeType, key string, path ...string)
	AddSetSymbol(name string, nodeType ast.NodeType) EntitySetSymbol
	AddFkSetSymbol(name string, linkedType ListStore) EntitySetSymbol

	NewRowComparator(sort []ast.SortField) (RowComparator, error)
	GetPublicSymbols() []string

	FindMatching(tx *bbolt.Tx, readIndex SetReadIndex, values []string) []string

	// QueryIds compiles the query and runs it against the store
	QueryIds(tx *bbolt.Tx, query string) ([][]byte, int64, error)

	// QueryIdsC executes a compile query against the store
	QueryIdsC(tx *bbolt.Tx, query ast.Query) ([][]byte, int64, error)
}

type CrudStore interface {
	ListStore

	GetParentStore() CrudStore
	AddLinkCollection(local EntitySymbol, remove EntitySymbol) LinkCollection
	GetLinkCollection(name string) LinkCollection

	Create(ctx MutateContext, entity BaseEntity) error
	Update(ctx MutateContext, entity BaseEntity, checker FieldChecker) error
	DeleteById(ctx MutateContext, id string) error
	DeleteWhere(ctx MutateContext, query string) error
	CleanupExternal(ctx MutateContext, id string) error

	CreateChild(ctx MutateContext, parentId string, entity BaseEntity) error
	UpdateChild(ctx MutateContext, parentId string, entity BaseEntity, checker FieldChecker) error
	DeleteChild(ctx MutateContext, parentId string, entity BaseEntity) error
	ListChildIds(tx *bbolt.Tx, parentId string, childType string) []string

	BaseLoadOneById(tx *bbolt.Tx, id string, entity BaseEntity) (bool, error)
	BaseLoadOneByQuery(tx *bbolt.Tx, query string, entity BaseEntity) (bool, error)
	BaseLoadOneChildById(tx *bbolt.Tx, id string, childId string, entity BaseEntity) (bool, error)
	NewStoreEntity() BaseEntity

	GetRelatedEntitiesIdList(tx *bbolt.Tx, id string, field string) []string

	AddDeleteHandler(handler EntityChangeHandler)
	events.EventEmmiter
}

type PersistContext struct {
	Id           string
	Store        CrudStore
	Bucket       *TypedBucket
	FieldChecker FieldChecker
	IsCreate     bool
}

func (ctx *PersistContext) GetParentContext() *PersistContext {
	return &PersistContext{
		Id:           ctx.Id,
		Store:        ctx.Store.GetParentStore(),
		Bucket:       ctx.Store.GetParentStore().GetEntityBucket(ctx.Bucket.Tx(), []byte(ctx.Id)),
		FieldChecker: ctx.FieldChecker,
		IsCreate:     ctx.IsCreate,
	}
}

func (ctx *PersistContext) WithFieldOverrides(overrides map[string]string) {
	if ctx.FieldChecker != nil {
		ctx.FieldChecker = NewMappedFieldChecker(ctx.FieldChecker, overrides)
	}
}

func (ctx *PersistContext) SetString(field string, value string) {
	ctx.Bucket.SetString(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetStringP(field string, value *string) {
	ctx.Bucket.SetStringP(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetTimeP(field string, value *time.Time) {
	ctx.Bucket.SetTimeP(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetBool(field string, value bool) {
	ctx.Bucket.SetBool(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetInt32(field string, value int32) {
	ctx.Bucket.SetInt32(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetMap(field string, value map[string]interface{}) {
	ctx.Bucket.PutMap(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetStringList(field string, value []string) {
	ctx.Bucket.SetStringList(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetLinkedIds(field string, value []string) {
	if ctx.Bucket.Err == nil && (ctx.FieldChecker == nil || ctx.FieldChecker.IsUpdated(field)) {
		serviceLinks := ctx.Store.GetLinkCollection(field)
		ctx.Bucket.SetError(serviceLinks.SetLinks(ctx.Bucket.Tx(), ctx.Id, value))
	}
}

type BaseEntity interface {
	GetId() string
	SetId(id string)
	LoadValues(store CrudStore, bucket *TypedBucket)
	SetValues(ctx *PersistContext)
	GetEntityType() string
}
