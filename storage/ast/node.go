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
	"time"
)

type NodeType int

const (
	NodeTypeBool NodeType = iota
	NodeTypeDatetime
	NodeTypeFloat64
	NodeTypeInt64
	NodeTypeString
	NodeTypeAnyType

	NodeTypeOther
)

func NodeTypeName(nodeType NodeType) string {
	return nodeTypeNames[nodeType]
}

var nodeTypeNames = map[NodeType]string{
	NodeTypeString:   "string",
	NodeTypeInt64:    "number",
	NodeTypeFloat64:  "number",
	NodeTypeDatetime: "date",
	NodeTypeBool:     "bool",
	NodeTypeOther:    "other",
	NodeTypeAnyType:  "any",
}

type BinaryOp int

const (
	BinaryOpEQ BinaryOp = iota
	BinaryOpNEQ
	BinaryOpLT
	BinaryOpLTE
	BinaryOpGT
	BinaryOpGTE
	BinaryOpIn
	BinaryOpNotIn
	BinaryOpBetween
	BinaryOpNotBetween
	BinaryOpContains
	BinaryOpNotContains
)

func (op BinaryOp) String() string {
	return binaryOpNames[op]
}

var binaryOpNames = map[BinaryOp]string{
	BinaryOpEQ:          "=",
	BinaryOpNEQ:         "!=",
	BinaryOpLT:          "<",
	BinaryOpLTE:         "<=",
	BinaryOpGT:          ">",
	BinaryOpGTE:         ">=",
	BinaryOpIn:          "in",
	BinaryOpNotIn:       "not in",
	BinaryOpBetween:     "between",
	BinaryOpNotBetween:  "not between",
	BinaryOpContains:    "contains",
	BinaryOpNotContains: "not contains",
}

var binaryOpValues = map[string]BinaryOp{
	"=":  BinaryOpEQ,
	"!=": BinaryOpNEQ,
	"<":  BinaryOpLT,
	"<=": BinaryOpLTE,
	">":  BinaryOpGT,
	">=": BinaryOpGTE,
}

type SetFunction int

const (
	SetFunctionAllOf SetFunction = iota
	SetFunctionAnyOf
	SetFunctionNoneOf
)

var SetFunctionNames = map[SetFunction]string{
	SetFunctionAllOf:  "allOf",
	SetFunctionAnyOf:  "anyOf",
	SetFunctionNoneOf: "noneOf",
}

type SymbolTypes interface {
	GetSymbolType(name string) (NodeType, bool)
	IsSet(name string) (bool, bool)
}

type Symbols interface {
	SymbolTypes
	EvalBool(name string) (*bool, error)
	EvalString(name string) (*string, error)
	EvalInt64(name string) (*int64, error)
	EvalFloat64(name string) (*float64, error)
	EvalDatetime(name string) (*time.Time, error)
	IsNil(name string) (bool, error)

	OpenSetCursor(name string) (SetCursor, error)
}

type SetCursor interface {
	Next()
	IsValid() bool
	Close()
}

type Node interface {
	fmt.Stringer
	GetType() NodeType
	Accept(visitor Visitor)
}

type TypeTransformable interface {
	Node
	TypeTransform(s SymbolTypes) (Node, error)
}

type BoolNode interface {
	Node
	EvalBool(s Symbols) (bool, error)
}

type BoolTypeTransformable interface {
	BoolNode
	TypeTransformBool(s SymbolTypes) (BoolNode, error)
}

type DatetimeNode interface {
	Node
	EvalDatetime(s Symbols) (*time.Time, error)
}

type Float64Node interface {
	StringNode
	EvalFloat64(s Symbols) (*float64, error)
}

type Int64Node interface {
	StringNode
	EvalInt64(s Symbols) (*int64, error)
	ToFloat64() Float64Node
}

type StringNode interface {
	Node
	EvalString(s Symbols) (*string, error)
}

type SymbolNode interface {
	Node
	Symbol() string
}

type AsStringArrayable interface {
	AsStringArray() *StringArrayNode
}

type SortField interface {
	Symbol() string
	IsAscending() bool
}

type Query interface {
	BoolNode

	// GetSortFields returns the fiels on which to sort. Returning nil or empty means the default sort order
	// will be used, usually by id ascending
	GetSortFields() []SortField

	// GetSkip returns the number of rows to skip. nil, or a values less than one will mean no rows skipped
	GetSkip() *int64

	// GetLimit returns the maximum number of rows to return. Returning nil will use the system configured
	// default for max rows. Returning -1 means do not limit results.
	GetLimit() *int64

	SetSkip(int64)
	SetLimit(int64)
}