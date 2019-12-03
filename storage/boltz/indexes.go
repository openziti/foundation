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

type fieldTypeAndValue struct {
	FieldType
	value []byte
}

type IndexingContext struct {
	isCreate   bool
	tx         *bbolt.Tx
	rowId      []byte
	errHolder  errorz.ErrorHolder
	atomStates map[Constraint][]byte
	setStates  map[Constraint][]fieldTypeAndValue
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
		isCreate:   isCreate,
		tx:         tx,
		rowId:      []byte(id),
		errHolder:  holder,
		atomStates: map[Constraint][]byte{},
		setStates:  map[Constraint][]fieldTypeAndValue{},
	}
}

func (indexer *Indexer) ProcessBeforeUpdate(ctx *IndexingContext) {
	if !ctx.errHolder.HasError() {
		for _, index := range indexer.constraints {
			index.processBeforeUpdate(ctx)
		}
	}
}

func (indexer *Indexer) ProcessAfterUpdate(ctx *IndexingContext) {
	if !ctx.errHolder.HasError() {
		for _, index := range indexer.constraints {
			index.processAfterUpdate(ctx)
		}
	}
}

func (indexer *Indexer) ProcessDelete(ctx *IndexingContext) {
	if !ctx.errHolder.HasError() {
		for _, index := range indexer.constraints {
			index.processDelete(ctx)
		}
	}
}

func (indexer *Indexer) InitializeIndexes(tx *bbolt.Tx, errorHolder errorz.ErrorHolder) {
	if !errorHolder.HasError() {
		for _, index := range indexer.constraints {
			index.initialize(tx, errorHolder)
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

type SetReadIndex interface {
	GetSymbol() EntitySetSymbol
	Read(tx *bbolt.Tx, key []byte, f func(val []byte))
}

type Constraint interface {
	processBeforeUpdate(ctx *IndexingContext)
	processAfterUpdate(ctx *IndexingContext)
	processDelete(ctx *IndexingContext)
	initialize(tx *bbolt.Tx, errorHolder errorz.ErrorHolder)
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

func (index *uniqueIndex) initialize(tx *bbolt.Tx, errorHolder errorz.ErrorHolder) {
	if !errorHolder.HasError() {
		pfxlog.Logger().Debugf("Ensuring bolt index bucket exists for %v", index.indexPath)
		bucket := GetOrCreatePath(tx, index.indexPath...)
		errorHolder.SetError(bucket.Err)
	}
}

func (index *uniqueIndex) processBeforeUpdate(ctx *IndexingContext) {
	if !ctx.errHolder.HasError() {
		_, fieldValue := index.symbol.Eval(ctx.tx, ctx.rowId)
		ctx.atomStates[index] = fieldValue
	}
}

func (index *uniqueIndex) processAfterUpdate(ctx *IndexingContext) {
	if !ctx.errHolder.HasError() {
		_, newValue := index.symbol.Eval(ctx.tx, ctx.rowId)
		oldValue := ctx.atomStates[index]

		if !ctx.isCreate && bytes.Equal(oldValue, newValue) {
			return
		}

		indexBucket := index.getIndexBucket(ctx.tx)

		if len(oldValue) > 0 {
			ctx.errHolder.SetError(indexBucket.DeleteValue(oldValue).Err)
		}

		if len(newValue) > 0 {
			ctx.errHolder.SetError(indexBucket.PutValue(newValue, ctx.rowId).Err)
		} else if !index.nullable {
			ctx.errHolder.SetError(errors.Errorf("index on %v.%v does not allow null or empty values",
				index.symbol.GetStore().GetEntityType(), index.symbol.GetName()))
		}
	}
}

func (index *uniqueIndex) processDelete(ctx *IndexingContext) {
	if !ctx.errHolder.HasError() {
		if _, value := index.symbol.Eval(ctx.tx, ctx.rowId); len(value) > 0 {
			indexBucket := index.getIndexBucket(ctx.tx)
			ctx.errHolder.SetError(indexBucket.DeleteValue(value).Err)
		}
	}
}

type setIndex struct {
	symbol    EntitySetSymbol
	indexPath []string
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
		_, value := getTypeAndValue(val)
		f(value)
	}
}

func (index *setIndex) visitCurrent(ctx *IndexingContext, f func(fieldType FieldType, value []byte)) {
	rtSymbol := index.symbol.GetRuntimeSymbol()
	cursor := rtSymbol.OpenCursor(ctx.tx, ctx.rowId)
	defer cursor.Close()
	for cursor.IsValid() {
		fieldType, value := rtSymbol.Eval(ctx.tx, ctx.rowId)
		f(fieldType, value)
		cursor.Next()
	}
}

func (index *setIndex) getCurrentValues(ctx *IndexingContext) []fieldTypeAndValue {
	var result []fieldTypeAndValue
	index.visitCurrent(ctx, func(fieldType FieldType, value []byte) {
		result = append(result, fieldTypeAndValue{
			FieldType: fieldType,
			value:     value,
		})
	})
	return result
}

func (index *setIndex) processBeforeUpdate(ctx *IndexingContext) {
	if !ctx.errHolder.HasError() {
		ctx.setStates[index] = index.getCurrentValues(ctx)
	}
}

func (index *setIndex) processAfterUpdate(ctx *IndexingContext) {
	if !ctx.errHolder.HasError() {
		oldValues := ctx.setStates[index]
		newValues := index.getCurrentValues(ctx)

		changed := false
		if len(oldValues) != len(newValues) {
			changed = true
		} else {
			for idx, oldVal := range oldValues {
				newVal := newValues[idx]
				if oldVal.FieldType != newVal.FieldType || !bytes.Equal(oldVal.value, newVal.value) {
					changed = true
					break
				}
			}
		}

		if !changed {
			return
		}

		for _, oldVal := range oldValues {
			indexBucket := index.getIndexBucket(ctx.tx, oldVal.value)
			ctx.errHolder.SetError(indexBucket.DeleteListEntry(TypeString, ctx.rowId).Err)
		}
		for _, newVal := range newValues {
			indexBucket := index.getIndexBucket(ctx.tx, newVal.value)
			ctx.errHolder.SetError(indexBucket.SetListEntry(TypeString, ctx.rowId).Err)
		}
	}
}

func (index *setIndex) processDelete(ctx *IndexingContext) {
	if !ctx.errHolder.HasError() {
		values := index.getCurrentValues(ctx)
		for _, val := range values {
			indexBucket := index.getIndexBucket(ctx.tx, val.value)
			ctx.errHolder.SetError(indexBucket.DeleteListEntry(TypeString, ctx.rowId).Err)
		}
	}
}

func (index *setIndex) initialize(tx *bbolt.Tx, errorHolder errorz.ErrorHolder) {
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

func (index *fkIndex) processBeforeUpdate(ctx *IndexingContext) {
	if !ctx.errHolder.HasError() {
		_, fieldValue := index.symbol.Eval(ctx.tx, ctx.rowId)
		ctx.atomStates[index] = fieldValue
	}
}

func (index *fkIndex) processAfterUpdate(ctx *IndexingContext) {
	if !ctx.errHolder.HasError() {
		_, newValue := index.symbol.Eval(ctx.tx, ctx.rowId)
		oldValue := ctx.atomStates[index]

		if !ctx.isCreate && bytes.Equal(oldValue, newValue) {
			return
		}

		if len(oldValue) > 0 {
			indexBucket := index.getIndexBucket(ctx.tx, oldValue)
			ctx.errHolder.SetError(indexBucket.DeleteListEntry(TypeString, ctx.rowId).Err)
		}

		if len(newValue) > 0 {
			indexBucket := index.getIndexBucket(ctx.tx, newValue)
			ctx.errHolder.SetError(indexBucket.SetListEntry(TypeString, ctx.rowId).Err)
		} else if !index.nullable {
			ctx.errHolder.SetError(errors.Errorf("index on %v.%v does not allow null or empty values",
				index.symbol.GetStore().GetEntityType(), index.symbol.GetName()))
		}
	}
}

func (index *fkIndex) processDelete(ctx *IndexingContext) {
	if !ctx.errHolder.HasError() {
		if _, value := index.symbol.Eval(ctx.tx, ctx.rowId); len(value) > 0 {
			indexBucket := index.getIndexBucket(ctx.tx, value)
			ctx.errHolder.SetError(indexBucket.DeleteListEntry(TypeString, ctx.rowId).Err)
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

func (index *fkIndex) initialize(_ *bbolt.Tx, _ errorz.ErrorHolder) {
	// nothing to do, as this index has no static location
}

type fkDeleteConstraint struct {
	symbol   EntitySetSymbol
	fkSymbol EntitySymbol
}

func (index *fkDeleteConstraint) processBeforeUpdate(_ *IndexingContext) {
}

func (index *fkDeleteConstraint) processAfterUpdate(_ *IndexingContext) {
}

func (index *fkDeleteConstraint) processDelete(ctx *IndexingContext) {
	if !ctx.errHolder.HasError() {
		rtSymbol := index.symbol.GetRuntimeSymbol()
		if rtSymbol.OpenCursor(ctx.tx, ctx.rowId).IsValid() {
			_, firstId := rtSymbol.Eval(ctx.tx, ctx.rowId)
			ctx.errHolder.SetError(errors.Errorf("cannot delete %v with id %v is referenced by %v with id %v, field %v",
				index.symbol.GetStore().GetEntityType(), string(ctx.rowId), index.fkSymbol.GetStore().GetEntityType(),
				string(firstId), index.fkSymbol.GetName()))
		}
	}
}

func (index *fkDeleteConstraint) initialize(_ *bbolt.Tx, _ errorz.ErrorHolder) {
	// nothing to do, as this index has no static location
}

