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
	"go.etcd.io/bbolt"
)

const (
	SortMax = 5
)

type RowComparator interface {
	Compare(rowId1, rowId2 RowCursor) int
}

type Scanner interface {
	Scan(tx *bbolt.Tx, query ast.Query) ([][]byte, int64, error)
}

type RowCursor interface {
	CurrentRow() []byte
	Tx() *bbolt.Tx
}

type EntitySymbol interface {
	GetStore() ListStore
	GetPath() []string
	GetType() ast.NodeType
	GetName() string
	IsSet() bool
	Eval(tx *bbolt.Tx, rowId []byte) (FieldType, []byte)
}

type EntitySetSymbol interface {
	EntitySymbol
	GetRuntimeSymbol() RuntimeEntitySetSymbol
	EvalStringList(tx *bbolt.Tx, key []byte) []string
}

type RuntimeEntitySetSymbol interface {
	EntitySymbol
	OpenCursor(tx *bbolt.Tx, rowId []byte) ast.SetCursor
}

type boltCursorFacade interface {
	Init()
	IsValid() bool
	Next()
	Id() []byte
}