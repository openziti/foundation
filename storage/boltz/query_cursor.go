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
	"time"

	"github.com/netfoundry/ziti-foundation/storage/ast"
	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

var _ ast.Symbols = (*rowCursorImpl)(nil)

type rowCursorImpl struct {
	symbolCache map[string]EntitySymbol
	entity      ListStore
	currentRow  []byte
	tx          *bbolt.Tx
}

func newRowCursor(entity ListStore, tx *bbolt.Tx) *rowCursorImpl {
	return &rowCursorImpl{
		symbolCache: map[string]EntitySymbol{},
		entity:      entity,
		tx:          tx,
	}
}

func (rs *rowCursorImpl) getSymbol(name string) EntitySymbol {
	result, found := rs.symbolCache[name]
	if !found {
		result = rs.entity.GetSymbol(name)
		if result != nil {
			rs.symbolCache[name] = result
		}
	}
	return result
}

func (rs *rowCursorImpl) GetSetSymbolTypes(name string) ast.SymbolTypes {
	return rs.entity.GetSetSymbolTypes(name)
}

func (rs *rowCursorImpl) NextRow(id []byte) {
	rs.currentRow = id
}

func (rs *rowCursorImpl) CurrentRow() []byte {
	return rs.currentRow
}

func (rs *rowCursorImpl) Tx() *bbolt.Tx {
	return rs.tx
}

func (rs *rowCursorImpl) GetSymbolType(name string) (ast.NodeType, bool) {
	symbol := rs.getSymbol(name)
	if symbol == nil {
		return 0, false
	}
	return symbol.GetType(), true
}

func (rs *rowCursorImpl) IsSet(name string) (bool, bool) {
	symbol := rs.getSymbol(name)
	if symbol == nil {
		return false, false
	}
	return symbol.IsSet(), true
}

func (rs *rowCursorImpl) OpenSetCursor(name string) (ast.SetCursor, error) {
	symbol := rs.getSymbol(name)
	if symbol == nil {
		return nil, errors.Errorf("unknown symbol %v", name)
	}
	setRowSymbol, ok := symbol.(RuntimeEntitySetSymbol)
	if !ok {
		return nil, errors.Errorf("attempting to iterate non-set symbol %v", name)
	}
	return setRowSymbol.OpenCursor(rs.tx, rs.currentRow), nil
}

func (rs *rowCursorImpl) OpenSetCursorForQuery(name string, query ast.Query) (ast.SetCursor, error) {
	// Using the shared symbol returned by rs.getSymbol(name) results in a race condition. I'm not sure how, because
	// everything seems to be either unshared or sequential. Further investigation warranted as time permits.
	symbol := rs.entity.GetSymbol(name)
	if symbol == nil {
		return nil, errors.Errorf("unknown symbol %v", name)
	}
	setRowSymbol, ok := symbol.(RuntimeEntitySetSymbol)
	if !ok {
		return nil, errors.Errorf("attempting to iterate non-set symbol %v", name)
	}
	setCursor := setRowSymbol.OpenCursor(rs.tx, rs.currentRow)
	return newCursorScanner(rs.tx, symbol.GetLinkedType(), setCursor, query)
}

func (rs *rowCursorImpl) EvalBool(name string) (*bool, error) {
	symbol := rs.getSymbol(name)
	if symbol == nil {
		return nil, errors.Errorf("unknown symbol %v", name)
	}
	return FieldToBool(symbol.Eval(rs.tx, rs.currentRow)), nil
}

func (rs *rowCursorImpl) EvalString(name string) (*string, error) {
	symbol := rs.getSymbol(name)
	if symbol == nil {
		return nil, errors.Errorf("unknown symbol %v", name)
	}
	return FieldToString(symbol.Eval(rs.tx, rs.currentRow)), nil
}

func (rs *rowCursorImpl) EvalInt64(name string) (*int64, error) {
	symbol := rs.getSymbol(name)
	if symbol == nil {
		return nil, errors.Errorf("unknown symbol %v", name)
	}
	return FieldToInt64(symbol.Eval(rs.tx, rs.currentRow)), nil
}

func (rs *rowCursorImpl) EvalFloat64(name string) (*float64, error) {
	symbol := rs.getSymbol(name)
	if symbol == nil {
		return nil, errors.Errorf("unknown symbol %v", name)
	}
	return FieldToFloat64(symbol.Eval(rs.tx, rs.currentRow)), nil
}

func (rs *rowCursorImpl) EvalDatetime(name string) (*time.Time, error) {
	symbol := rs.getSymbol(name)
	if symbol == nil {
		return nil, errors.Errorf("unknown symbol %v", name)
	}
	fieldType, val := symbol.Eval(rs.tx, rs.currentRow)
	return FieldToDatetime(fieldType, val, symbol.GetName()), nil
}

func (rs *rowCursorImpl) IsNil(name string) (bool, error) {
	symbol := rs.getSymbol(name)
	if symbol == nil {
		return false, errors.Errorf("unknown symbol %v", name)
	}
	fieldType, _ := symbol.Eval(rs.tx, rs.currentRow)
	return fieldType == TypeNil, nil
}
