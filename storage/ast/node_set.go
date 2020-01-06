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

package ast

import (
	"fmt"
	"strconv"
)

type AllOfSetExprNode struct {
	name      string
	predicate BoolNode
}

func (node *AllOfSetExprNode) Symbol() string {
	return node.name
}

func (node *AllOfSetExprNode) String() string {
	return fmt.Sprintf("allOf(%v)", node.predicate)
}

func (node *AllOfSetExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *AllOfSetExprNode) Accept(visitor Visitor) {
	visitor.VisitAllOfSetExprNodeStart(node)
	node.predicate.Accept(visitor)
	visitor.VisitAllOfSetExprNodeEnd(node)
}

func (node *AllOfSetExprNode) EvalBool(s Symbols) (result bool, err error) {
	var cursor SetCursor
	result = true

	cursor, err = s.OpenSetCursor(node.name)
	if err != nil || cursor == nil {
		return
	}

	for cursor.IsValid() {
		result, err = node.predicate.EvalBool(s)
		if err != nil {
			return false, err
		}
		if !result {
			return
		}
		if err := cursor.Next(); err != nil {
			return false, err
		}
	}
	return
}

type AnyOfSetExprNode struct {
	name      string
	predicate BoolNode
}

func (node *AnyOfSetExprNode) Symbol() string {
	return node.name
}

func (node *AnyOfSetExprNode) String() string {
	return fmt.Sprintf("anyOf(%v)", node.predicate)
}

func (node *AnyOfSetExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *AnyOfSetExprNode) Accept(visitor Visitor) {
	visitor.VisitAnyOfSetExprNodeStart(node)
	node.predicate.Accept(visitor)
	visitor.VisitAnyOfSetExprNodeEnd(node)
}

func (node *AnyOfSetExprNode) EvalBool(s Symbols) (bool, error) {
	cursor, err := s.OpenSetCursor(node.name)
	if err != nil || cursor == nil {
		return false, err
	}

	for cursor.IsValid() {
		result, err := node.predicate.EvalBool(s)
		if err != nil {
			return false, err
		}
		if result {
			return true, nil
		}
		if err := cursor.Next(); err != nil {
			return false, err
		}
	}
	return false, nil
}

type CountSetExprNode struct {
	symbol SymbolNode
	query  Query
}

func (node *CountSetExprNode) Symbol() string {
	return node.symbol.Symbol()
}

func (node *CountSetExprNode) String() string {
	return fmt.Sprintf("count(%v)", node.Symbol())
}

func (node *CountSetExprNode) GetType() NodeType {
	return NodeTypeInt64
}

func (node *CountSetExprNode) Accept(visitor Visitor) {
	visitor.VisitCountSetExprNodeStart(node)
	node.symbol.Accept(visitor)
	if node.query != nil {
		node.query.Accept(visitor)
	}
	visitor.VisitCountSetExprNodeEnd(node)
}

func (node *CountSetExprNode) EvalInt64(s Symbols) (*int64, error) {
	var result int64
	var cursor SetCursor
	var err error

	if node.query == nil {
		cursor, err = s.OpenSetCursor(node.Symbol())
	} else {
		cursor, err = s.OpenSetCursorForQuery(node.Symbol(), node.query)
	}

	if err != nil {
		return nil, err
	}

	if cursor == nil {
		return &result, err
	}

	for cursor.IsValid() {
		result++
		if err := cursor.Next(); err != nil {
			return nil, err
		}
	}
	return &result, nil
}

func (node *CountSetExprNode) EvalString(s Symbols) (*string, error) {
	result, err := node.EvalInt64(s)
	if err != nil {
		return nil, err
	}
	stringResult := strconv.FormatInt(*result, 10)
	return &stringResult, nil
}

func (node *CountSetExprNode) ToFloat64() Float64Node {
	return &Int64ToFloat64Node{node}
}

type IsEmptySetExprNode struct {
	symbol SymbolNode
	query  Query
}

func (node *IsEmptySetExprNode) Symbol() string {
	return node.symbol.Symbol()
}

func (node *IsEmptySetExprNode) String() string {
	return fmt.Sprintf("isEmpty(%v)", node.Symbol())
}

func (node *IsEmptySetExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *IsEmptySetExprNode) Accept(visitor Visitor) {
	visitor.VisitIsEmptySetExprNodeStart(node)
	node.symbol.Accept(visitor)
	if node.query != nil {
		node.query.Accept(visitor)
	}
	visitor.VisitIsEmptySetExprNodeEnd(node)
}

func (node *IsEmptySetExprNode) EvalBool(s Symbols) (bool, error) {
	var cursor SetCursor
	var err error

	if node.query == nil {
		cursor, err = s.OpenSetCursor(node.Symbol())
	} else {
		cursor, err = s.OpenSetCursorForQuery(node.Symbol(), node.query)
	}
	if err != nil {
		return true, err
	}
	if cursor == nil {
		return true, err
	}
	return !cursor.IsValid(), nil
}
