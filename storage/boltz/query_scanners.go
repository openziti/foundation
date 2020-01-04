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
	"math"

	"github.com/biogo/store/llrb"
	"github.com/netfoundry/ziti-foundation/storage/ast"
	"go.etcd.io/bbolt"
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

	cursor    ast.SetCursor
	rowCursor *rowCursorImpl
	filter    ast.BoolNode
	current   []byte
}

func newCursorScanner(tx *bbolt.Tx, store ListStore, cursor ast.SetCursor, query ast.Query) (ast.SetCursor, error) {
	result := &uniqueIndexScanner{
		store:     store,
		forward:   true,
		cursor:    cursor,
		rowCursor: newRowCursor(store, tx),
		filter:    query,
	}
	result.setPaging(query)
	if err := result.Next(); err != nil {
		return nil, err
	}
	return result, nil
}

func (scanner *uniqueIndexScanner) Scan(tx *bbolt.Tx, query ast.Query) ([]string, int64, error) {
	scanner.setPaging(query)
	entityBucket := scanner.store.GetEntitiesBucket(tx)
	if entityBucket == nil {
		return nil, 0, nil
	}
	boltCursor := entityBucket.Cursor()
	scanner.rowCursor = newRowCursor(scanner.store, boltCursor.Bucket().Tx())
	scanner.filter = query
	scanner.cursor = NewBoltCursor(boltCursor, scanner.forward)
	if err := scanner.Next(); err != nil {
		return nil, 0, err
	}

	var result []string
	for scanner.IsValid() {
		id := scanner.Current()
		if scanner.offset < scanner.targetOffset {
			scanner.offset++
		} else {
			if scanner.count < scanner.targetLimit {
				result = append(result, string(id))
			}
			scanner.count++
		}
		if err := scanner.Next(); err != nil {
			return nil, 0, err
		}
	}
	return result, scanner.count, nil
}

func (scanner *uniqueIndexScanner) IsValid() bool {
	return scanner.current != nil
}

func (scanner *uniqueIndexScanner) Current() []byte {
	return scanner.current
}

func (scanner *uniqueIndexScanner) Next() error {
	cursor := scanner.cursor
	rowCursor := scanner.rowCursor
	for {
		if !cursor.IsValid() {
			scanner.current = nil
			return nil
		}

		scanner.current = cursor.Current()
		if err := cursor.Next(); err != nil {
			scanner.current = nil
			return err
		}
		if scanner.store.IsChildStore() && !scanner.store.IsEntityPresent(rowCursor.Tx(), string(scanner.current)) {
			continue
		}
		rowCursor.NextRow(scanner.current)
		match, err := scanner.filter.EvalBool(rowCursor)
		if err != nil {
			return err
		}
		if match {
			return nil
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
	cursor := NewForwardBoltCursor(entityBucket.Cursor())
	rowContext := &RowContext{comparator: comparator, rowCursor1: rowCursor, rowCursor2: newRowCursor(scanner.store, tx)}

	// Longer term, if we're looking for better performance, we could make a version of llrb which takes a comparator
	// function instead of putting the comparison on the elements, so we don't need to start a context with each row
	results := &llrb.Tree{}
	isChildStore := scanner.store.IsChildStore()
	maxResults := scanner.targetOffset + scanner.targetLimit
	for cursor.IsValid() {
		current := cursor.Current()
		if err := cursor.Next(); err != nil {
			return nil, 0, err
		}
		if isChildStore && !scanner.store.IsEntityPresent(tx, string(current)) {
			continue
		}
		rowCursor.NextRow(current)
		match, err := query.EvalBool(rowCursor)
		if err != nil {
			return nil, 0, err
		}
		if match {
			results.Insert(&Row{id: current, context: rowContext})
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
