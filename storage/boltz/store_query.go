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
	"github.com/michaelquigley/pfxlog"
	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
	"strings"
)

func (store *BaseStore) GetPublicSymbols() []string {
	return store.publicSymbols
}

func (store *BaseStore) GetSymbolType(name string) (ast.NodeType, bool) {
	if symbol := store.GetSymbol(name); symbol != nil {
		return symbol.GetType(), true
	}
	return 0, false
}

func (store *BaseStore) IsSet(name string) (bool, bool) {
	if symbol := store.GetSymbol(name); symbol != nil {
		return symbol.IsSet(), true
	}
	return false, false
}

// GetSymbol returns the symbol for the given name, or nil if the symbol doesn't exist
func (store *BaseStore) GetSymbol(name string) EntitySymbol {
	/*
		Types of symbols that we need to handle
		1. Local single values (employee.name)
		2. Local set values (sub-buckets of non-id keys) (myEntity.phoneNumbers)
		3. Composite single value symbols (employee.manager.name)
		4. Composite multi-value symbols (employee.directReports.phoneNumbers)
		5. Maps (employee.tags.location, employee.manager.tags.location, employee.directReports.tags.location)
	*/
	if result := store.symbols[name]; result != nil {
		// If it's a set symbol, make a runtime copy so we don't share cursor data. If we ever have a case where we
		// are evaluating the same symbol in multiple context, this will still break, but since currently any given
		// expression only involves a single symbol, this should not be a problem
		if setSymbol, ok := result.(EntitySetSymbol); ok {
			return setSymbol.GetRuntimeSymbol()
		}
		return result
	}

	// if it's a composite symbol, create a symbol on the fly to represent the name
	if index := strings.IndexRune(name, '.'); index > 0 {
		parts := strings.SplitN(name, ".", 2)
		// If it's a map symbol, create that now and return it
		if mapSymbol := store.mapSymbols[parts[0]]; mapSymbol != nil {
			var prefix []string
			prefix = append(prefix, mapSymbol.prefix...)
			prefix = append(prefix, mapSymbol.key)
			return store.newEntitySymbol(name, mapSymbol.symbolType, parts[1], nil, prefix...)
		}

		if result := store.GetSymbol(parts[0]); result != nil {
			linkedEntitySymbol, ok := result.(linkedEntitySymbol)
			if !ok || linkedEntitySymbol.getLinkedType() == nil {
				return nil // Can only have composite symbols if it's linked
			}
			rest := linkedEntitySymbol.getLinkedType().GetSymbol(parts[1])
			if rest == nil {
				return nil
			}
			return store.createCompositeEntitySymbol(name, linkedEntitySymbol, rest)
		}
	}

	return nil
}

func (store *BaseStore) createCompositeEntitySymbol(name string, first linkedEntitySymbol, rest EntitySymbol) EntitySymbol {
	ces, ok := rest.(compositeEntitySymbol)
	var chain []EntitySymbol
	if !ok {
		chain = []EntitySymbol{first, rest}
	} else {
		chain = []EntitySymbol{first}
		chain = append(chain, ces.getChain()...)
	}

	noneSet := true
	var last EntitySymbol
	var iterable []iterableEntitySymbol
	for _, symbol := range chain {
		if symbol.IsSet() {
			noneSet = false
		}
		if iterableSymbol, ok := symbol.(iterableEntitySymbol); ok {
			iterable = append(iterable, iterableSymbol)
		}
		last = symbol
	}
	// strip ids, since they are redundant
	if _, ok := last.(*entityIdSymbol); ok {
		chain = chain[:len(chain)-1]
		if len(chain) == 1 {
			return chain[0]
		}
		last = chain[len(chain)-1]
	}
	if noneSet {
		return &nonSetCompositeEntitySymbol{
			name:       name,
			symbolType: rest.GetType(),
			chain:      chain,
		}
	}
	if len(chain) == len(iterable) {
		return &compositeEntitySetSymbol{
			name:       name,
			symbolType: rest.GetType(),
			chain:      iterable,
			cursor:     nil,
			cursorLastF: func(tx *bbolt.Tx, key []byte) (fieldType FieldType, bytes []byte) {
				return getTypeAndValue(key)
			},
		}
	}
	return &compositeEntitySetSymbol{
		name:        name,
		symbolType:  rest.GetType(),
		chain:       iterable,
		cursor:      nil,
		cursorLastF: last.Eval,
	}
}

func (store *BaseStore) addPublicSymbol(name string, symbol EntitySymbol) EntitySymbol {
	store.symbols[name] = symbol
	store.publicSymbols = append(store.publicSymbols, name)
	return symbol
}

func (store *BaseStore) GrantSymbols(child ListStore) {
	for _, value := range store.symbols {
		child.inheritSymbol(value)
	}
}

func (store *BaseStore) inheritSymbol(symbol EntitySymbol) {
	store.symbols[symbol.GetName()] = symbol
}

func (store *BaseStore) AddIdSymbol(name string, nodeType ast.NodeType) EntitySymbol {
	return store.addPublicSymbol(name, &entityIdSymbol{
		store:      store,
		symbolType: nodeType,
		path:       []string{"id"},
	})
}

func (store *BaseStore) AddSymbol(name string, nodeType ast.NodeType, prefix ...string) EntitySymbol {
	return store.AddSymbolWithKey(name, nodeType, name, prefix...)
}

func (store *BaseStore) AddFkSymbol(name string, linkedStore ListStore, prefix ...string) EntitySymbol {
	return store.AddFkSymbolWithKey(name, name, linkedStore, prefix...)
}
func (store *BaseStore) AddSymbolWithKey(name string, nodeType ast.NodeType, key string, prefix ...string) EntitySymbol {
	pfxlog.Logger().Debugf("adding symbol %v.%v [entity]", store.GetEntityType(), name)
	return store.addPublicSymbol(name, store.newEntitySymbol(name, nodeType, key, nil, prefix...))
}

func (store *BaseStore) AddFkSymbolWithKey(name string, key string, linkedStore ListStore, prefix ...string) EntitySymbol {
	pfxlog.Logger().Debugf("adding symbol %v.%v [entity]", store.GetEntityType(), name)
	return store.addPublicSymbol(name, store.newEntitySymbol(name, ast.NodeTypeString, key, linkedStore, prefix...))
}

func (store *BaseStore) AddMapSymbol(name string, nodeType ast.NodeType, key string, prefix ...string) {
	pfxlog.Logger().Debugf("adding symbol %v.%v [map]", store.GetEntityType(), name)
	store.mapSymbols[name] = &entityMapSymbol{
		key:        key,
		symbolType: nodeType,
		prefix:     prefix,
	}
}

func (store *BaseStore) newEntitySymbol(name string, nodeType ast.NodeType, key string, linkedType ListStore, prefix ...string) *entitySymbol {
	var path []string
	var bucketF func(entityBucket *TypedBucket) *TypedBucket

	if len(prefix) == 0 {
		path = []string{key}
		bucketF = func(entityBucket *TypedBucket) *TypedBucket {
			return entityBucket
		}
	} else {
		path = append(path, prefix...)
		path = append(path, key)

		bucketF = func(entityBucket *TypedBucket) *TypedBucket {
			if entityBucket == nil {
				return nil
			}
			return entityBucket.GetPath(prefix...)
		}
	}
	return &entitySymbol{
		store:      store,
		name:       name,
		getBucketF: bucketF,
		symbolType: nodeType,
		prefix:     prefix,
		key:        key,
		path:       path,
		linkedType: linkedType,
	}
}

func (store *BaseStore) AddSetSymbol(name string, nodeType ast.NodeType) EntitySetSymbol {
	return store.addSetSymbol(name, nodeType, nil)
}

func (store *BaseStore) AddFkSetSymbol(name string, listStore ListStore) EntitySetSymbol {
	return store.addSetSymbol(name, ast.NodeTypeString, listStore)
}

func (store *BaseStore) addSetSymbol(name string, nodeType ast.NodeType, listStore ListStore) EntitySetSymbol {
	pfxlog.Logger().Debugf("adding symbol %v.%v [set]", store.GetEntityType(), name)
	entitySymbol := store.newEntitySymbol(name, nodeType, name, listStore)
	result := &entitySetSymbolImpl{
		entitySymbol: entitySymbol,
	}
	store.symbols[name] = result
	return result
}

func (store *BaseStore) NewScanner(sort []ast.SortField) (Scanner, error) {
	if len(sort) > SortMax {
		sort = sort[:SortMax]
	}

	if len(sort) == 0 || sort[0].Symbol() == "id" {
		if len(sort) < 1 || sort[0].IsAscending() {
			return &uniqueIndexScanner{store: store, forward: true}, nil
		}
		return &uniqueIndexScanner{store: store, forward: false}, nil
	}
	return &sortingScanner{
		store: store,
	}, nil
}



type sortFieldImpl struct {
	name  string
	isAsc bool
}

func (sortField *sortFieldImpl) Symbol() string {
	return sortField.name
}

func (sortField *sortFieldImpl) IsAscending() bool {
	return sortField.isAsc
}

func (store *BaseStore) NewRowComparator(sort []ast.SortField) (RowComparator, error) {
	if len(sort) == 0 {
		sort = []ast.SortField{&sortFieldImpl{name: "id", isAsc: true}}
	}

	var symbolsComparators []symbolComparator
	for _, sortField := range sort {
		symbol, found := store.symbols[sortField.Symbol()]
		forward := sortField.IsAscending()
		if !found {
			return nil, errors.Errorf("no such sort field: %v", sortField.Symbol())
		}
		if symbol.IsSet() {
			return nil, errors.Errorf("invalid sort field: %v", sortField.Symbol())
		}

		var symbolComparator symbolComparator
		switch symbol.GetType() {
		case ast.NodeTypeBool:
			symbolComparator = &boolSymbolComparator{symbol: symbol, forward: forward}
		case ast.NodeTypeDatetime:
			symbolComparator = &datetimeSymbolComparator{symbol: symbol, forward: forward}
		case ast.NodeTypeFloat64:
			symbolComparator = &float64SymbolComparator{symbol: symbol, forward: forward}
		case ast.NodeTypeInt64:
			symbolComparator = &int64SymbolComparator{symbol: symbol, forward: forward}
		case ast.NodeTypeString:
			symbolComparator = &stringSymbolComparator{symbol: symbol, forward: forward}
		default:
			return nil, errors.Errorf("unsupported sort field type %v for field : %v", ast.NodeTypeName(symbol.GetType()), sortField.Symbol())
		}
		symbolsComparators = append(symbolsComparators, symbolComparator)
	}

	return &rowComparatorImpl{symbols: symbolsComparators}, nil
}

func (store *BaseStore) QueryIds(tx *bbolt.Tx, queryString string) ([][]byte, int64, error) {
	query, err := ast.Parse(store, queryString)
	if err != nil {
		return nil, 0, err
	}
	return store.QueryIdsC(tx, query)
}

func (store *BaseStore) QueryIdsC(tx *bbolt.Tx, query ast.Query) ([][]byte, int64, error) {
	scanner, err := store.NewScanner(query.GetSortFields())
	if err != nil {
		return nil, 0, err
	}
	return scanner.Scan(tx, query)
}