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
	"strings"
)

type QueryNodeImpl struct {
	Predicate BoolNode
	SortBy    *SortByNode
	Skip      *SkipExprNode
	Limit     *LimitExprNode
}

func (node *QueryNodeImpl) SetSkip(skip int64) {
	node.Skip = &SkipExprNode{Int64ConstNode{value: skip}}
}

func (node *QueryNodeImpl) SetLimit(limit int64) {
	node.Limit = &LimitExprNode{Int64ConstNode{value: limit}}
}

func (node *QueryNodeImpl) TypeTransformBool(s SymbolTypes) (BoolNode, error) {
	if err := transformBools(s, &node.Predicate); err != nil {
		return node, err
	}
	if _, err := node.SortBy.TypeTransform(s); err != nil {
		return node, err
	}
	return node, nil
}

func (node *QueryNodeImpl) EvalBool(s Symbols) (bool, error) {
	return node.Predicate.EvalBool(s)
}

func (node *QueryNodeImpl) GetPredicate() BoolNode {
	return node.Predicate
}

func (node *QueryNodeImpl) GetSortFields() []SortField {
	return node.SortBy.getSortFields()
}

func (node *QueryNodeImpl) GetSkip() *int64 {
	if node.Skip == nil {
		return nil
	}
	return &node.Skip.value
}

func (node *QueryNodeImpl) GetLimit() *int64 {
	if node.Limit == nil {
		return nil
	}
	return &node.Limit.value
}

func (node *QueryNodeImpl) String() string {
	return node.Predicate.String() + " " + node.SortBy.String()
}

func (node *QueryNodeImpl) GetType() NodeType {
	return NodeTypeOther
}

func (node *QueryNodeImpl) Accept(visitor Visitor) {
	node.Predicate.Accept(visitor)
	node.SortBy.Accept(visitor)
	node.Skip.Accept(visitor)
	node.Limit.Accept(visitor)
	visitor.VisitQueryNode(node)
}

type SortByNode struct {
	SortFields []*SortFieldNode
}

func (node *SortByNode) getSortFields() []SortField {
	if node == nil {
		return nil
	}
	var result []SortField
	for _, sortField := range node.SortFields {
		result = append(result, sortField)
	}
	return result
}

func (node *SortByNode) TypeTransform(s SymbolTypes) (Node, error) {
	if node != nil {
		for _, sortField := range node.SortFields {
			if _, err := sortField.TypeTransform(s); err != nil {
				return node, err
			}
		}
	}
	return node, nil
}

func (node *SortByNode) String() string {
	builder := strings.Builder{}
	builder.WriteString("sort by ")
	if node == nil || len(node.SortFields) == 0 {
		builder.WriteString("<default>")
	} else {
		builder.WriteString(node.SortFields[0].String())
		for _, sortField := range node.SortFields[1:] {
			builder.WriteString(", ")
			builder.WriteString(sortField.String())
		}
	}
	return builder.String()
}

func (node *SortByNode) GetType() NodeType {
	return NodeTypeOther
}

func (node *SortByNode) Accept(visitor Visitor) {
	if node != nil {
		for _, sortField := range node.SortFields {
			sortField.Accept(visitor)
		}
		visitor.VisitSortByNode(node)
	}
}

type SortFieldNode struct {
	symbol      SymbolNode
	isAscending bool
}

func (node *SortFieldNode) Symbol() string {
	return node.symbol.Symbol()
}

func (node *SortFieldNode) IsAscending() bool {
	return node.isAscending
}

func (node *SortFieldNode) TypeTransform(s SymbolTypes) (Node, error) {
	var symbolNode Node = node.symbol
	err := transformTypes(s, &symbolNode)
	node.symbol = symbolNode.(SymbolNode)
	return node, err
}

func (node *SortFieldNode) String() string {
	if node.isAscending {
		return fmt.Sprintf("%v ASC", node.symbol)
	}
	return fmt.Sprintf("%v DESC", node.symbol)
}

func (node *SortFieldNode) GetType() NodeType {
	return NodeTypeOther
}

func (node *SortFieldNode) Accept(visitor Visitor) {
	node.symbol.Accept(visitor)
	visitor.VisitSortFieldNode(node)
}

type SortDirection bool

const (
	SortAscending  SortDirection = true
	SortDescending SortDirection = false
)

type LimitExprNode struct {
	Int64ConstNode
}

func (node *LimitExprNode) String() string {
	if node.value == -1 {
		return "limit none"
	}
	return fmt.Sprintf("limit %v", node.value)
}

func (node *LimitExprNode) Accept(visitor Visitor) {
	if node != nil {
		visitor.VisitLimitExprNode(node)
	}
}

type SkipExprNode struct {
	Int64ConstNode
}

func (node *SkipExprNode) String() string {
	return fmt.Sprintf("skip %v", node.value)
}

func (node *SkipExprNode) Accept(visitor Visitor) {
	if node != nil {
		visitor.VisitSkipExprNode(node)
	}
}


