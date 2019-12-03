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
	"github.com/pkg/errors"
	"strconv"
	"time"
)

func NewUntypedSymbolNode(symbol string) SymbolNode {
	return &UntypedSymbolNode{symbol: symbol}
}

type UntypedSymbolNode struct {
	symbol string
}

func (node *UntypedSymbolNode) Accept(visitor Visitor) {
	visitor.VisitSymbol(node.symbol, node.GetType())
	visitor.VisitUntypedSymbolNode(node)
}

func (node *UntypedSymbolNode) TypeTransform(s SymbolTypes) (Node, error) {
	kind, found := s.GetSymbolType(node.symbol)
	if !found {
		return nil, errors.Errorf("unknown symbol %v", node.symbol)
	}
	switch kind {
	case NodeTypeString:
		return &StringSymbolNode{symbol: node.symbol}, nil
	case NodeTypeBool:
		return &BoolSymbolNode{symbol: node.symbol}, nil
	case NodeTypeInt64:
		return &Int64SymbolNode{symbol: node.symbol}, nil
	case NodeTypeFloat64:
		return &Float64SymbolNode{symbol: node.symbol}, nil
	case NodeTypeDatetime:
		return &DatetimeSymbolNode{symbol: node.symbol}, nil
	case NodeTypeAnyType:
		return &AnyTypeSymbolNode{symbol: node.symbol}, nil
	}
	return nil, errors.Errorf("unhanded symbol type %v for symbol %v", kind, node.symbol)
}

func (node *UntypedSymbolNode) GetType() NodeType {
	return NodeTypeOther
}

func (node *UntypedSymbolNode) String() string {
	return node.symbol
}

func (node *UntypedSymbolNode) Symbol() string {
	return node.symbol
}

// BoolSymbolNode implements lookup of symbol values of type bool
type BoolSymbolNode struct {
	symbol string
}

func (node *BoolSymbolNode) Accept(visitor Visitor) {
	visitor.VisitSymbol(node.symbol, node.GetType())
	visitor.VisitBoolSymbolNode(node)
}

func (node *BoolSymbolNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *BoolSymbolNode) EvalBool(s Symbols) (bool, error) {
	result, err := s.EvalBool(node.symbol)
	if err != nil {
		return false, err
	}
	return result != nil && *result, nil
}

func (node *BoolSymbolNode) String() string {
	return node.symbol
}

func (node *BoolSymbolNode) Symbol() string {
	return node.symbol
}

// DatetimeSymbolNode implements lookup of symbol values of type datetime
type DatetimeSymbolNode struct {
	symbol string
}

func (node *DatetimeSymbolNode) Accept(visitor Visitor) {
	visitor.VisitSymbol(node.symbol, node.GetType())
	visitor.VisitDatetimeSymbolNode(node)
}

func (node *DatetimeSymbolNode) GetType() NodeType {
	return NodeTypeDatetime
}

func (node *DatetimeSymbolNode) EvalDatetime(s Symbols) (*time.Time, error) {
	return s.EvalDatetime(node.symbol)
}

func (node *DatetimeSymbolNode) String() string {
	return node.symbol
}

func (node *DatetimeSymbolNode) Symbol() string {
	return node.symbol
}

// Float64SymbolNode implements lookup of symbol values of type float64
type Float64SymbolNode struct {
	symbol string
}

func (node *Float64SymbolNode) GetType() NodeType {
	return NodeTypeFloat64
}

func (node *Float64SymbolNode) Accept(visitor Visitor) {
	visitor.VisitSymbol(node.symbol, node.GetType())
	visitor.VisitFloat64SymbolNode(node)
}

func (node *Float64SymbolNode) EvalFloat64(s Symbols) (*float64, error) {
	return s.EvalFloat64(node.symbol)
}

func (node *Float64SymbolNode) EvalString(s Symbols) (*string, error) {
	float64Val, err := s.EvalFloat64(node.symbol)
	if err != nil {
		return nil, err
	}
	if float64Val != nil {
		result := strconv.FormatFloat(*float64Val, 'f', -1, 64)
		return &result, nil
	}
	return nil, nil
}

func (node *Float64SymbolNode) String() string {
	return node.symbol
}

func (node *Float64SymbolNode) Symbol() string {
	return node.symbol
}

// Int64SymbolNode implements lookup of symbol values of type int64
type Int64SymbolNode struct {
	symbol string
}

func (node *Int64SymbolNode) GetType() NodeType {
	return NodeTypeInt64
}

func (node *Int64SymbolNode) Accept(visitor Visitor) {
	visitor.VisitSymbol(node.symbol, node.GetType())
	visitor.VisitInt64SymbolNode(node)
}

func (node *Int64SymbolNode) EvalInt64(s Symbols) (*int64, error) {
	return s.EvalInt64(node.symbol)
}

func (node *Int64SymbolNode) EvalString(s Symbols) (*string, error) {
	int64Val, err := s.EvalInt64(node.symbol)
	if err != nil {
		return nil, err
	}
	if int64Val != nil {
		result := strconv.FormatInt(*int64Val, 10)
		return &result, nil
	}
	return nil, nil
}

func (node *Int64SymbolNode) ToFloat64() Float64Node {
	return &Int64ToFloat64Node{node}
}

func (node *Int64SymbolNode) String() string {
	return node.symbol
}

func (node *Int64SymbolNode) Symbol() string {
	return node.symbol
}

// StringSymbolNode implements lookup of symbol values of type string
type StringSymbolNode struct {
	symbol string
}

func (node *StringSymbolNode) GetType() NodeType {
	return NodeTypeString
}

func (node *StringSymbolNode) Accept(visitor Visitor) {
	visitor.VisitSymbol(node.symbol, node.GetType())
	visitor.VisitStringSymbolNode(node)
}

func (node *StringSymbolNode) EvalString(s Symbols) (*string, error) {
	return s.EvalString(node.symbol)
}

func (node *StringSymbolNode) String() string {
	return node.symbol
}

func (node *StringSymbolNode) Symbol() string {
	return node.symbol
}

// AnyTypeSymbolNode implements lookup of symbol values of any, meaning they can have any value type
type AnyTypeSymbolNode struct {
	symbol string
}

func (node *AnyTypeSymbolNode) GetType() NodeType {
	return NodeTypeAnyType
}

func (node *AnyTypeSymbolNode) Accept(visitor Visitor) {
	visitor.VisitSymbol(node.symbol, node.GetType())
	visitor.VisitAnyTypeSymbolNode(node)
}

func (node *AnyTypeSymbolNode) EvalBool(s Symbols) (bool, error) {
	result, err := s.EvalBool(node.symbol)
	if err != nil {
		return false, err
	}
	return result != nil && *result, nil
}

func (node *AnyTypeSymbolNode) EvalDatetime(s Symbols) (*time.Time, error) {
	return s.EvalDatetime(node.symbol)
}

func (node *AnyTypeSymbolNode) EvalInt64(s Symbols) (*int64, error) {
	return s.EvalInt64(node.symbol)
}

func (node *AnyTypeSymbolNode) ToFloat64() Float64Node {
	return node
}

func (node *AnyTypeSymbolNode) EvalFloat64(s Symbols) (*float64, error) {
	return s.EvalFloat64(node.symbol)
}

func (node *AnyTypeSymbolNode) EvalString(s Symbols) (*string, error) {
	return s.EvalString(node.symbol)
}

func (node *AnyTypeSymbolNode) String() string {
	return node.symbol
}

func (node *AnyTypeSymbolNode) Symbol() string {
	return node.symbol
}

type SymbolValidator struct {
	DefaultVisitor
	inSetFunction bool
	symbolTypes   SymbolTypes
	err           error
}

func (visitor *SymbolValidator) VisitSetFunctionNodeStart(node *SetFunctionNode) {
	visitor.inSetFunction = true
}

func (visitor *SymbolValidator) VisitSetFunctionNodeEnd(node *SetFunctionNode) {
	visitor.inSetFunction = false

	isSet, found := visitor.symbolTypes.IsSet(node.symbol.Symbol())
	if found && !isSet {
		visitor.err = fmt.Errorf("symbol '%v' is not a set symbol but is used in set function %v", node.symbol.Symbol(), SetFunctionNames[node.setFunction])
	}
}

func (visitor *SymbolValidator) VisitUntypedSymbolNode(node *UntypedSymbolNode) {
	isSet, found := visitor.symbolTypes.IsSet(node.Symbol())
	if !found {
		visitor.err = fmt.Errorf("unknown symbol %v", node.Symbol())
		return
	}

	if !visitor.inSetFunction && isSet {
		visitor.err = fmt.Errorf("symbol '%v' is a set symbol but is used in non-set function context", node.Symbol())
	}
}
