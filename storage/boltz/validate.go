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
	"github.com/netfoundry/ziti-foundation/util/stringz"
	"github.com/pkg/errors"
)

type publicSymbolValidator struct {
	ast.DefaultVisitor
	store ListStore
	err   error
}

func (visitor *publicSymbolValidator) VisitSymbol(symbol string, _ ast.NodeType) {
	if visitor.err == nil && !stringz.Contains(visitor.store.GetPublicSymbols(), symbol) {
		visitor.err = errors.Errorf("invalid query identifier '%v' for type %v", symbol, visitor.store.GetEntityType())
	}
}

func ValidateSymbolsArePublic(query ast.Query, store ListStore) error {
	visitor := &publicSymbolValidator{store: store}
	query.Accept(visitor)
	return visitor.err
}
