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
	"github.com/biogo/store/llrb"
	"go.etcd.io/bbolt"
	"math"
)

type scanner struct {
	targetOffset int64
	targetLimit  int64
}

func (s *scanner) setPaging(query ast.Query) {
	if query.GetSkip() == nil {
		query.SetSkip(0)
	}
	s.targetOffset = *query.GetSkip()

	if query.GetLimit() == nil {
		query.SetLimit(math.MaxInt64)
	}

	if *query.GetLimit() < 0 {
		query.SetLimit(math.MaxInt64)
	}

	s.targetLimit = *query.GetLimit()
}

type uniqueIndexScanner struct {
	scanner
	store   ListStore
	forward bool
	offset  int64
	count   int64
}

func (scanner *uniqueIndexScanner) Scan(tx *bbolt.Tx, query ast.Query) ([][]byte, int64, error) {
	scanner.setPaging(query)
	entityBucket := scanner.store.GetEntitiesBucket(tx)
	if entityBucket == nil {
		return nil, 0, nil
	}
	boltCursor := entityBucket.Cursor()

	rowCursor := newRowCursor(scanner.store, tx)

	var cursor boltCursorFacade
	if scanner.forward {
		cursor = &ForwardBoltCursor{BaseBoltCursor{cursor: boltCursor}}
	} else {
		cursor = &ReverseBoltCursor{BaseBoltCursor{cursor: boltCursor}}
	}

	var result [][]byte
	isChildStore := scanner.store.IsChildStore()
	for cursor.Init(); cursor.IsValid(); cursor.Next() {
		id := cursor.Id()
		if isChildStore && !scanner.store.IsEntityPresent(tx, string(id)) {
			continue
		}
		rowCursor.NextRow(id)
		match, err := query.EvalBool(rowCursor)
		if err != nil {
			return nil, 0, err
		}
		if match {
			if scanner.offset < scanner.targetOffset {
				scanner.offset++
			} else {
				if scanner.count < scanner.targetLimit {
					result = append(result, id)
				}
				scanner.count++
			}
		}
	}

	return result, scanner.count, nil
}

type sortingScanner struct {
	scanner
	store  ListStore
	offset int64
	count  int64
}

func (scanner *sortingScanner) Scan(tx *bbolt.Tx, query ast.Query) ([][]byte, int64, error) {
	scanner.setPaging(query)
	comparator, err := scanner.store.NewRowComparator(query.GetSortFields())
	if err != nil {
		return nil, 0, err
	}
	entityBucket := scanner.store.GetEntitiesBucket(tx)
	if entityBucket == nil {
		return nil, 0, nil
	}

	rowCursor := newRowCursor(scanner.store, tx)
	cursor := &ForwardBoltCursor{BaseBoltCursor{cursor: entityBucket.Cursor()}}
	rowContext := &RowContext{comparator: comparator, rowCursor1: rowCursor, rowCursor2: newRowCursor(scanner.store, tx)}

	// Longer term, if we're looking for better performance, we could make a version of llrb which takes a comparator
	// function instead of putting the comparison on the elements, so we don't need to start a context with each row
	results := &llrb.Tree{}
	isChildStore := scanner.store.IsChildStore()
	maxResults := scanner.targetOffset + scanner.targetLimit
	for cursor.Init(); cursor.IsValid(); cursor.Next() {
		if isChildStore && !scanner.store.IsEntityPresent(tx, string(cursor.Id())) {
			continue
		}
		rowCursor.NextRow(cursor.Id())
		match, err := query.EvalBool(rowCursor)
		if err != nil {
			return nil, 0, err
		}
		if match {
			results.Insert(&Row{id: cursor.Id(), context: rowContext})
			scanner.count++
			if scanner.count > maxResults {
				results.DeleteMax()
			}
		}
	}

	var result [][]byte
	results.Do(func(row llrb.Comparable) bool {
		if scanner.offset < scanner.targetOffset {
			scanner.offset++
		} else {
			result = append(result, row.(*Row).id)
		}
		return false
	})

	return result, scanner.count, nil
}

type RowContext struct {
	comparator RowComparator
	rowCursor1 *rowCursorImpl
	rowCursor2 *rowCursorImpl
}

type Row struct {
	id      []byte
	context *RowContext
}

func (r *Row) Compare(other llrb.Comparable) int {
	r.context.rowCursor1.NextRow(r.id)
	r.context.rowCursor2.NextRow(other.(*Row).id)
	return r.context.comparator.Compare(r.context.rowCursor1, r.context.rowCursor2)
}
