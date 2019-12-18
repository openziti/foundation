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
	"github.com/biogo/store/llrb"
	"github.com/netfoundry/ziti-foundation/storage/ast"
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

	cursor    boltCursorFacade
	rowCursor *rowCursorImpl
	filter    ast.BoolNode
}

func (scanner *uniqueIndexScanner) Scan(tx *bbolt.Tx, query ast.Query) ([]string, int64, error) {
	scanner.setPaging(query)
	entityBucket := scanner.store.GetEntitiesBucket(tx)
	if entityBucket == nil {
		return nil, 0, nil
	}
	boltCursor := entityBucket.Cursor()
	scanner.init(boltCursor, query)

	var result []string
	for {
		id, valid, err := scanner.next()
		if err != nil {
			return nil, 0, err
		}
		if valid {
			if scanner.offset < scanner.targetOffset {
				scanner.offset++
			} else {
				if scanner.count < scanner.targetLimit {
					result = append(result, string(id))
				}
				scanner.count++
			}
		} else {
			return result, scanner.count, nil
		}
	}
}

func (scanner *uniqueIndexScanner) init(boltCursor *bbolt.Cursor, filter ast.BoolNode) {
	scanner.rowCursor = newRowCursor(scanner.store, boltCursor.Bucket().Tx())
	scanner.filter = filter

	if scanner.forward {
		scanner.cursor = &ForwardBoltCursor{BaseBoltCursor{cursor: boltCursor}}
	} else {
		scanner.cursor = &ReverseBoltCursor{BaseBoltCursor{cursor: boltCursor}}
	}
	scanner.cursor.Init()
}

func (scanner *uniqueIndexScanner) next() ([]byte, bool, error) {
	cursor := scanner.cursor
	rowCursor := scanner.rowCursor
	for {
		if !cursor.IsValid() {
			return nil, false, nil
		}

		id := cursor.Id()
		cursor.Next()
		if scanner.store.IsChildStore() && !scanner.store.IsEntityPresent(rowCursor.Tx(), string(id)) {
			continue
		}
		rowCursor.NextRow(id)
		match, err := scanner.filter.EvalBool(rowCursor)
		if err != nil {
			return nil, false, err
		}
		if match {
			return id, true, nil
		}
	}
}

type sortingScanner struct {
	scanner
	store  ListStore
	offset int64
	count  int64
}

func (scanner *sortingScanner) Scan(tx *bbolt.Tx, query ast.Query) ([]string, int64, error) {
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

	var result []string
	results.Do(func(row llrb.Comparable) bool {
		if scanner.offset < scanner.targetOffset {
			scanner.offset++
		} else {
			result = append(result, string(row.(*Row).id))
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
