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

package boltz

import (
	"bytes"

	"github.com/michaelquigley/pfxlog"
	"github.com/netfoundry/ziti-foundation/util/errorz"
	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

/*

/ziti/services/<id>/fabric-properties
/ziti/services/<id>/edge/edge-properties
/ziti/services/edge/<index-name>

*/

type Indexer struct {
	constraints []Constraint
	basePath    []string
}

type FieldTypeAndValue struct {
	FieldType
	Value []byte
}

type IndexingContext struct {
	IsCreate   bool
	Tx         *bbolt.Tx
	RowId      []byte
	ErrHolder  errorz.ErrorHolder
	atomStates map[Constraint][]byte
	setStates  map[Constraint][]FieldTypeAndValue
}

func NewIndexer(basePath ...string) *Indexer {
	return &Indexer{basePath: basePath}
}

func (indexer *Indexer) AddUniqueIndex(symbol EntitySymbol) ReadIndex {
	index := &uniqueIndex{
		symbol:    symbol,
		nullable:  false,
		indexPath: indexer.getIndexPath(symbol),
	}

	indexer.constraints = append(indexer.constraints, index)
	return index
}

func (indexer *Indexer) AddSetIndex(symbol EntitySetSymbol) SetReadIndex {
	index := &setIndex{
		symbol:    symbol,
		indexPath: indexer.getIndexPath(symbol),
	}
	indexer.constraints = append(indexer.constraints, index)
	return index
}

func (indexer *Indexer) AddFkIndex(symbol EntitySymbol, fkSymbol EntitySetSymbol) {
	indexer.addFkIndex(symbol, fkSymbol, false)
}

func (indexer *Indexer) AddNullableFkIndex(symbol EntitySymbol, fkSymbol EntitySetSymbol) {
	indexer.addFkIndex(symbol, fkSymbol, true)
}

func (indexer *Indexer) addFkIndex(symbol EntitySymbol, fkSymbol EntitySetSymbol, nullable bool) {
	index := &fkIndex{
		symbol:   symbol,
		nullable: nullable,
		fkSymbol: fkSymbol,
	}

	indexer.addConstraint(index)
	fkStore := fkSymbol.GetStore()
	if baseStore, ok := fkStore.(*BaseStore); ok {
		baseStore.addConstraint(&fkDeleteConstraint{
			symbol:   fkSymbol,
			fkSymbol: symbol,
		})
	} else {
		pfxlog.Logger().Warnf("fk store %v is not an indexer, can't enforce validity of constraint on delete",
			fkSymbol.GetStore().GetEntityType())
	}
}

func (indexer *Indexer) addConstraint(constraint Constraint) {
	indexer.constraints = append(indexer.constraints, constraint)
}

func (indexer *Indexer) NewIndexingContext(isCreate bool, tx *bbolt.Tx, id string, holder errorz.ErrorHolder) *IndexingContext {
	return &IndexingContext{
		IsCreate:   isCreate,
		Tx:         tx,
		RowId:      []byte(id),
		ErrHolder:  holder,
		atomStates: map[Constraint][]byte{},
		setStates:  map[Constraint][]FieldTypeAndValue{},
	}
}

func (indexer *Indexer) ProcessBeforeUpdate(ctx *IndexingContext) {
	if !ctx.ErrHolder.HasError() {
		for _, index := range indexer.constraints {
			index.ProcessBeforeUpdate(ctx)
		}
	}
}

func (indexer *Indexer) ProcessAfterUpdate(ctx *IndexingContext) {
	if !ctx.ErrHolder.HasError() {
		for _, index := range indexer.constraints {
			index.ProcessAfterUpdate(ctx)
		}
	}
}

func (indexer *Indexer) ProcessDelete(ctx *IndexingContext) {
	if !ctx.ErrHolder.HasError() {
		for _, index := range indexer.constraints {
			index.ProcessDelete(ctx)
		}
	}
}

func (indexer *Indexer) InitializeIndexes(tx *bbolt.Tx, errorHolder errorz.ErrorHolder) {
	if !errorHolder.HasError() {
		for _, index := range indexer.constraints {
			index.Initialize(tx, errorHolder)
		}
	}
}

func (indexer *Indexer) getIndexPath(symbol EntitySymbol) []string {
	var result []string
	result = append(result, indexer.basePath...)
	result = append(result, symbol.GetStore().GetEntityType(), symbol.GetName())
	return result
}

type ReadIndex interface {
	Read(tx *bbolt.Tx, val []byte) []byte
}

type SetChangeListener func(tx *bbolt.Tx, rowId []byte, old []FieldTypeAndValue, new []FieldTypeAndValue, holder errorz.ErrorHolder)

type SetReadIndex interface {
	GetSymbol() EntitySetSymbol
	Read(tx *bbolt.Tx, key []byte, f func(val []byte))
	AddListener(listener SetChangeListener)
}

type Constraint interface {
	ProcessBeforeUpdate(ctx *IndexingContext)
	ProcessAfterUpdate(ctx *IndexingContext)
	ProcessDelete(ctx *IndexingContext)
	Initialize(tx *bbolt.Tx, errorHolder errorz.ErrorHolder)
}

type uniqueIndex struct {
	symbol    EntitySymbol
	nullable  bool
	indexPath []string
}

func (index *uniqueIndex) Read(tx *bbolt.Tx, val []byte) []byte {
	indexBucket := index.getIndexBucket(tx)
	if indexBucket.Err != nil {
		return nil
	}
	return indexBucket.Get(val)
}

func (index *uniqueIndex) getIndexBucket(tx *bbolt.Tx) *TypedBucket {
	indexBucket := Path(tx, index.indexPath...)
	if indexBucket != nil {
		return indexBucket
	}
	return GetOrCreatePath(tx, index.indexPath...)
}

func (index *uniqueIndex) Initialize(tx *bbolt.Tx, errorHolder errorz.ErrorHolder) {
	if !errorHolder.HasError() {
		pfxlog.Logger().Debugf("Ensuring bolt index bucket exists for %v", index.indexPath)
		bucket := GetOrCreatePath(tx, index.indexPath...)
		errorHolder.SetError(bucket.Err)
	}
}

func (index *uniqueIndex) ProcessBeforeUpdate(ctx *IndexingContext) {
	if !ctx.ErrHolder.HasError() {
		_, fieldValue := index.symbol.Eval(ctx.Tx, ctx.RowId)
		ctx.atomStates[index] = fieldValue
	}
}

func (index *uniqueIndex) ProcessAfterUpdate(ctx *IndexingContext) {
	if !ctx.ErrHolder.HasError() {
		_, newValue := index.symbol.Eval(ctx.Tx, ctx.RowId)
		oldValue := ctx.atomStates[index]

		if !ctx.IsCreate && bytes.Equal(oldValue, newValue) {
			return
		}

		indexBucket := index.getIndexBucket(ctx.Tx)

		if len(oldValue) > 0 {
			ctx.ErrHolder.SetError(indexBucket.DeleteValue(oldValue).Err)
		}

		if len(newValue) > 0 {
			if indexBucket.Get(newValue) != nil {
				ctx.ErrHolder.SetError(errors.Errorf("duplicate value '%v' in unique index on %v store",
					string(newValue), index.symbol.GetStore().GetEntityType()))
			} else {
				ctx.ErrHolder.SetError(indexBucket.PutValue(newValue, ctx.RowId).Err)
			}
		} else if !index.nullable {
			ctx.ErrHolder.SetError(errors.Errorf("index on %v.%v does not allow null or empty values",
				index.symbol.GetStore().GetEntityType(), index.symbol.GetName()))
		}
	}
}

func (index *uniqueIndex) ProcessDelete(ctx *IndexingContext) {
	if !ctx.ErrHolder.HasError() {
		if _, value := index.symbol.Eval(ctx.Tx, ctx.RowId); len(value) > 0 {
			indexBucket := index.getIndexBucket(ctx.Tx)
			ctx.ErrHolder.SetError(indexBucket.DeleteValue(value).Err)
		}
	}
}

type setIndex struct {
	symbol    EntitySetSymbol
	indexPath []string
	listeners []SetChangeListener
}

func (index *setIndex) AddListener(listener SetChangeListener) {
	index.listeners = append(index.listeners, listener)
}

func (index *setIndex) GetSymbol() EntitySetSymbol {
	return index.symbol
}

func (index *setIndex) Read(tx *bbolt.Tx, key []byte, f func(val []byte)) {
	indexBaseBucket := Path(tx, index.indexPath...)
	if indexBaseBucket == nil {
		return
	}
	indexBucket := indexBaseBucket.Bucket.Bucket(key)
	if indexBucket == nil {
		return
	}
	cursor := indexBucket.Cursor()
	for val, _ := cursor.First(); val != nil; val, _ = cursor.Next() {
		_, value := GetTypeAndValue(val)
		f(value)
	}
}

func (index *setIndex) visitCurrent(ctx *IndexingContext, f func(fieldType FieldType, value []byte)) {
	rtSymbol := index.symbol.GetRuntimeSymbol()
	cursor := rtSymbol.OpenCursor(ctx.Tx, ctx.RowId)
	for cursor.IsValid() {
		fieldType, value := rtSymbol.Eval(ctx.Tx, ctx.RowId)
		f(fieldType, value)
		if err := cursor.Next(); err != nil {
			ctx.ErrHolder.SetError(err)
			return
		}
	}
}

func (index *setIndex) getCurrentValues(ctx *IndexingContext) []FieldTypeAndValue {
	var result []FieldTypeAndValue
	index.visitCurrent(ctx, func(fieldType FieldType, value []byte) {
		result = append(result, FieldTypeAndValue{
			FieldType: fieldType,
			Value:     value,
		})
	})
	return result
}

func (index *setIndex) ProcessBeforeUpdate(ctx *IndexingContext) {
	if !ctx.ErrHolder.HasError() {
		ctx.setStates[index] = index.getCurrentValues(ctx)
	}
}

func (index *setIndex) ProcessAfterUpdate(ctx *IndexingContext) {
	if !ctx.ErrHolder.HasError() {
		oldValues := ctx.setStates[index]
		newValues := index.getCurrentValues(ctx)

		changed := false
		if len(oldValues) != len(newValues) {
			changed = true
		} else {
			for idx, oldVal := range oldValues {
				newVal := newValues[idx]
				if oldVal.FieldType != newVal.FieldType || !bytes.Equal(oldVal.Value, newVal.Value) {
					changed = true
					break
				}
			}
		}

		if !changed {
			return
		}

		for _, oldVal := range oldValues {
			indexBucket := index.getIndexBucket(ctx.Tx, oldVal.Value)
			ctx.ErrHolder.SetError(indexBucket.DeleteListEntry(TypeString, ctx.RowId).Err)
		}
		for _, newVal := range newValues {
			indexBucket := index.getIndexBucket(ctx.Tx, newVal.Value)
			ctx.ErrHolder.SetError(indexBucket.SetListEntry(TypeString, ctx.RowId).Err)
		}
		for _, listener := range index.listeners {
			listener(ctx.Tx, ctx.RowId, oldValues, newValues, ctx.ErrHolder)
		}
	}
}

func (index *setIndex) ProcessDelete(ctx *IndexingContext) {
	if !ctx.ErrHolder.HasError() {
		values := index.getCurrentValues(ctx)
		for _, val := range values {
			indexBucket := index.getIndexBucket(ctx.Tx, val.Value)
			ctx.ErrHolder.SetError(indexBucket.DeleteListEntry(TypeString, ctx.RowId).Err)
		}
	}
}

func (index *setIndex) Initialize(tx *bbolt.Tx, errorHolder errorz.ErrorHolder) {
	if !errorHolder.HasError() {
		pfxlog.Logger().Debugf("Ensuring bolt index bucket exists for %v", index.indexPath)
		bucket := GetOrCreatePath(tx, index.indexPath...)
		errorHolder.SetError(bucket.Err)
	}
}

func (index *setIndex) getIndexBucket(tx *bbolt.Tx, key []byte) *TypedBucket {
	indexBucket := Path(tx, index.indexPath...)
	if indexBucket == nil {
		return ErrBucket(errors.Errorf("bucket at %+v for index not created", index.indexPath))
	}
	return indexBucket.GetOrCreateBucket(string(key))
}

type fkIndex struct {
	symbol   EntitySymbol
	fkSymbol EntitySymbol
	nullable bool
}

func (index *fkIndex) ProcessBeforeUpdate(ctx *IndexingContext) {
	if !ctx.ErrHolder.HasError() {
		_, fieldValue := index.symbol.Eval(ctx.Tx, ctx.RowId)
		ctx.atomStates[index] = fieldValue
	}
}

func (index *fkIndex) ProcessAfterUpdate(ctx *IndexingContext) {
	if !ctx.ErrHolder.HasError() {
		_, newValue := index.symbol.Eval(ctx.Tx, ctx.RowId)
		oldValue := ctx.atomStates[index]

		if !ctx.IsCreate && bytes.Equal(oldValue, newValue) {
			return
		}

		if len(oldValue) > 0 {
			indexBucket := index.getIndexBucket(ctx.Tx, oldValue)
			ctx.ErrHolder.SetError(indexBucket.DeleteListEntry(TypeString, ctx.RowId).Err)
		}

		if len(newValue) > 0 {
			indexBucket := index.getIndexBucket(ctx.Tx, newValue)
			ctx.ErrHolder.SetError(indexBucket.SetListEntry(TypeString, ctx.RowId).Err)
		} else if !index.nullable {
			ctx.ErrHolder.SetError(errors.Errorf("index on %v.%v does not allow null or empty values",
				index.symbol.GetStore().GetEntityType(), index.symbol.GetName()))
		}
	}
}

func (index *fkIndex) ProcessDelete(ctx *IndexingContext) {
	if !ctx.ErrHolder.HasError() {
		if _, value := index.symbol.Eval(ctx.Tx, ctx.RowId); len(value) > 0 {
			indexBucket := index.getIndexBucket(ctx.Tx, value)
			ctx.ErrHolder.SetError(indexBucket.DeleteListEntry(TypeString, ctx.RowId).Err)
		}
	}
}

func (index *fkIndex) getIndexBucket(tx *bbolt.Tx, fkId []byte) *TypedBucket {
	fkStore := index.fkSymbol.GetStore()
	entityBucket := fkStore.GetEntityBucket(tx, fkId)
	if entityBucket == nil {
		return ErrBucket(errors.Errorf("no entity of type %v with id %v", fkStore.GetEntityType(), string(fkId)))
	}
	return entityBucket.GetOrCreatePath(index.fkSymbol.GetPath()...)
}

func (index *fkIndex) Initialize(_ *bbolt.Tx, _ errorz.ErrorHolder) {
	// nothing to do, as this index has no static location
}

type fkDeleteConstraint struct {
	symbol   EntitySetSymbol
	fkSymbol EntitySymbol
}

func (index *fkDeleteConstraint) ProcessBeforeUpdate(_ *IndexingContext) {
}

func (index *fkDeleteConstraint) ProcessAfterUpdate(_ *IndexingContext) {
}

func (index *fkDeleteConstraint) ProcessDelete(ctx *IndexingContext) {
	if !ctx.ErrHolder.HasError() {
		rtSymbol := index.symbol.GetRuntimeSymbol()
		if rtSymbol.OpenCursor(ctx.Tx, ctx.RowId).IsValid() {
			_, firstId := rtSymbol.Eval(ctx.Tx, ctx.RowId)
			ctx.ErrHolder.SetError(errors.Errorf("cannot delete %v with id %v is referenced by %v with id %v, field %v",
				index.symbol.GetStore().GetEntityType(), string(ctx.RowId), index.fkSymbol.GetStore().GetEntityType(),
				string(firstId), index.fkSymbol.GetName()))
		}
	}
}

func (index *fkDeleteConstraint) Initialize(_ *bbolt.Tx, _ errorz.ErrorHolder) {
	// nothing to do, as this index has no static location
}
