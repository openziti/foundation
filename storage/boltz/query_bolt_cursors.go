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

package boltz

import (
	"github.com/netfoundry/ziti-foundation/storage/ast"
	"go.etcd.io/bbolt"
)

type BaseBoltCursor struct {
	cursor *bbolt.Cursor
	key    []byte
}

func (f *BaseBoltCursor) IsValid() bool {
	return f.key != nil
}

func (f *BaseBoltCursor) Current() []byte {
	return f.key
}

func NewBoltCursor(cursor *bbolt.Cursor, forward bool) ast.SetCursor {
	if forward {
		return NewForwardBoltCursor(cursor)
	}
	return NewReverseBoltCursor(cursor)
}

func NewForwardBoltCursor(cursor *bbolt.Cursor) ast.SetCursor {
	result := &ForwardBoltCursor{BaseBoltCursor{
		cursor: cursor,
		key:    nil,
	}}
	result.key, _ = result.cursor.First()
	return result
}

type ForwardBoltCursor struct {
	BaseBoltCursor
}

func (f *ForwardBoltCursor) Next() error {
	f.key, _ = f.cursor.Next()
	return nil
}

func NewReverseBoltCursor(cursor *bbolt.Cursor) ast.SetCursor {
	result := &ReverseBoltCursor{BaseBoltCursor{
		cursor: cursor,
		key:    nil,
	}}
	result.key, _ = result.cursor.Last()
	return result
}

type ReverseBoltCursor struct {
	BaseBoltCursor
}

func (f *ReverseBoltCursor) Next() error {
	f.key, _ = f.cursor.Prev()
	return nil
}

func NewTypedForwardBoltCursor(cursor *bbolt.Cursor) ast.SetCursor {
	result := &TypedForwardBoltCursor{BaseBoltCursor{
		cursor: cursor,
		key:    nil,
	}}

	key, _ := result.cursor.First()
	_, result.key = GetTypeAndValue(key)

	return result
}

type TypedForwardBoltCursor struct {
	BaseBoltCursor
}

func (f *TypedForwardBoltCursor) Next() error {
	key, _ := f.cursor.Next()
	_, f.key = GetTypeAndValue(key)
	return nil
}

func NewTypedReverseBoltCursor(cursor *bbolt.Cursor) ast.SetCursor {
	result := &TypedReverseBoltCursor{BaseBoltCursor{
		cursor: cursor,
		key:    nil,
	}}

	key, _ := result.cursor.Last()
	_, result.key = GetTypeAndValue(key)

	return result
}

type TypedReverseBoltCursor struct {
	BaseBoltCursor
}

func (f *TypedReverseBoltCursor) Next() error {
	key, _ := f.cursor.Prev()
	_, f.key = GetTypeAndValue(key)
	return nil
}
