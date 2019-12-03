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
	"strings"
)

// NotExprNode implements logical NOT on a wrapped boolean expression
type NotExprNode struct {
	node BoolNode
}

func (node *NotExprNode) Accept(visitor Visitor) {
	visitor.VisitNotExprNodeStart(node)
	node.node.Accept(visitor)
	visitor.VisitNotExprNodeEnd(node)
}

func (node *NotExprNode) String() string {
	return fmt.Sprintf("not (%v)", node.node)
}

func (node *NotExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *NotExprNode) EvalBool(s Symbols) (bool, error) {
	result, err := node.node.EvalBool(s)
	if err != nil {
		return false, err
	}
	return !result, nil
}

func (node *NotExprNode) TypeTransformBool(s SymbolTypes) (BoolNode, error) {
	return node, transformBools(s, &node.node)
}

// AndExprNode implements logical AND on two wrapped boolean expressions
type AndExprNode struct {
	left  BoolNode
	right BoolNode
}

func (node *AndExprNode) Accept(visitor Visitor) {
	visitor.VisitAndExprNodeStart(node)
	node.left.Accept(visitor)
	node.right.Accept(visitor)
	visitor.VisitAndExprNodeEnd(node)
}

func (node *AndExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *AndExprNode) TypeTransformBool(s SymbolTypes) (BoolNode, error) {
	return node, transformBools(s, &node.left, &node.right)
}

func (node *AndExprNode) EvalBool(s Symbols) (bool, error) {
	leftResult, err := node.left.EvalBool(s)
	if err != nil {
		return false, err
	}

	if !leftResult {
		return false, nil
	}

	return node.right.EvalBool(s)
}

func (node *AndExprNode) String() string {
	return fmt.Sprintf("%v && %v", node.left, node.right)
}

// OrExprNode implements logical OR on two wrapped boolean expressions
type OrExprNode struct {
	left  BoolNode
	right BoolNode
}

func (node *OrExprNode) Accept(visitor Visitor) {
	visitor.VisitOrExprNodeStart(node)
	node.left.Accept(visitor)
	node.right.Accept(visitor)
	visitor.VisitOrExprNodeEnd(node)
}

func (node *OrExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *OrExprNode) TypeTransformBool(s SymbolTypes) (BoolNode, error) {
	return node, transformBools(s, &node.left, &node.right)
}

func (node *OrExprNode) EvalBool(s Symbols) (bool, error) {
	leftResult, err := node.left.EvalBool(s)
	if err != nil {
		return false, err
	}
	if leftResult {
		return true, nil
	}
	return node.right.EvalBool(s)
}

func (node *OrExprNode) String() string {
	return fmt.Sprintf("%v || %v", node.left, node.right)
}

type BinaryBoolExprNode struct {
	left  BoolNode
	right BoolNode
	op    BinaryOp
}

func (node *BinaryBoolExprNode) Accept(visitor Visitor) {
	visitor.VisitBinaryBoolExprNodeStart(node)
	node.left.Accept(visitor)
	node.right.Accept(visitor)
	visitor.VisitBinaryBoolExprNodeEnd(node)
}

func (*BinaryBoolExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *BinaryBoolExprNode) EvalBool(s Symbols) (bool, error) {
	leftResult, err := node.left.EvalBool(s)
	if err != nil {
		return false, err
	}
	rightResult, err := node.right.EvalBool(s)
	if err != nil {
		return false, err
	}

	switch node.op {
	case BinaryOpEQ:
		return leftResult == rightResult, nil
	case BinaryOpNEQ:
		return leftResult != rightResult, nil
	}

	return false, errors.Errorf("unhandled boolean binary expression type %v", node.op)
}

func (node *BinaryBoolExprNode) String() string {
	return fmt.Sprintf("%v %v %v", node.left, binaryOpNames[node.op], node.right)
}

type BinaryDatetimeExprNode struct {
	left  DatetimeNode
	right DatetimeNode
	op    BinaryOp
}

func (node *BinaryDatetimeExprNode) Accept(visitor Visitor) {
	visitor.VisitBinaryDatetimeExprNodeStart(node)
	node.left.Accept(visitor)
	node.right.Accept(visitor)
	visitor.VisitBinaryDatetimeExprNodeEnd(node)
}

func (*BinaryDatetimeExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *BinaryDatetimeExprNode) EvalBool(s Symbols) (bool, error) {
	leftResult, err := node.left.EvalDatetime(s)
	if err != nil {
		return false, err
	}
	rightResult, err := node.right.EvalDatetime(s)
	if err != nil {
		return false, err
	}

	if leftResult == nil || rightResult == nil {
		return false, nil
	}

	switch node.op {
	case BinaryOpEQ:
		return leftResult.Equal(*rightResult), nil
	case BinaryOpNEQ:
		return !leftResult.Equal(*rightResult), nil
	case BinaryOpLT:
		return leftResult.Before(*rightResult), nil
	case BinaryOpLTE:
		return !leftResult.After(*rightResult), nil
	case BinaryOpGT:
		return leftResult.After(*rightResult), nil
	case BinaryOpGTE:
		return !leftResult.Before(*rightResult), nil
	}

	return false, errors.Errorf("unhandled datetime binary expression type %v", node.op)
}

func (node *BinaryDatetimeExprNode) String() string {
	return fmt.Sprintf("%v %v %v", node.left, binaryOpNames[node.op], node.right)
}

type BinaryFloat64ExprNode struct {
	left  Float64Node
	right Float64Node
	op    BinaryOp
}

func (node *BinaryFloat64ExprNode) Accept(visitor Visitor) {
	visitor.VisitBinaryFloat64ExprNodeStart(node)
	node.left.Accept(visitor)
	node.right.Accept(visitor)
	visitor.VisitBinaryFloat64ExprNodeEnd(node)
}

func (node *BinaryFloat64ExprNode) GetType() NodeType {
	return NodeTypeFloat64
}

func (node *BinaryFloat64ExprNode) EvalBool(s Symbols) (bool, error) {
	leftResult, err := node.left.EvalFloat64(s)
	if err != nil {
		return false, err
	}
	rightResult, err := node.right.EvalFloat64(s)
	if err != nil {
		return false, err
	}

	if leftResult == nil || rightResult == nil {
		return false, nil
	}

	switch node.op {
	case BinaryOpEQ:
		return *leftResult == *rightResult, nil
	case BinaryOpNEQ:
		return *leftResult != *rightResult, nil
	case BinaryOpLT:
		return *leftResult < *rightResult, nil
	case BinaryOpLTE:
		return *leftResult <= *rightResult, nil
	case BinaryOpGT:
		return *leftResult > *rightResult, nil
	case BinaryOpGTE:
		return *leftResult >= *rightResult, nil
	}

	return false, errors.Errorf("unhandled float64 binary expression type %v", node.op)
}

func (node *BinaryFloat64ExprNode) String() string {
	return fmt.Sprintf("%v %v %v", node.left, binaryOpNames[node.op], node.right)
}

type BinaryInt64ExprNode struct {
	left  Int64Node
	right Int64Node
	op    BinaryOp
}

func (node *BinaryInt64ExprNode) Accept(visitor Visitor) {
	visitor.VisitBinaryInt64ExprNodeStart(node)
	node.left.Accept(visitor)
	node.right.Accept(visitor)
	visitor.VisitBinaryInt64ExprNodeEnd(node)
}

func (node *BinaryInt64ExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *BinaryInt64ExprNode) EvalBool(s Symbols) (bool, error) {
	leftResult, err := node.left.EvalInt64(s)
	if err != nil {
		return false, err
	}
	rightResult, err := node.right.EvalInt64(s)
	if err != nil {
		return false, err
	}

	if leftResult == nil || rightResult == nil {
		return false, nil
	}

	switch node.op {
	case BinaryOpEQ:
		return *leftResult == *rightResult, nil
	case BinaryOpNEQ:
		return *leftResult != *rightResult, nil
	case BinaryOpLT:
		return *leftResult < *rightResult, nil
	case BinaryOpLTE:
		return *leftResult <= *rightResult, nil
	case BinaryOpGT:
		return *leftResult > *rightResult, nil
	case BinaryOpGTE:
		return *leftResult >= *rightResult, nil
	}

	return false, errors.Errorf("unhandled int64 binary expression type %v", node.op)
}

func (node *BinaryInt64ExprNode) String() string {
	return fmt.Sprintf("%v %v %v", node.left, binaryOpNames[node.op], node.right)
}

type BinaryStringExprNode struct {
	left  StringNode
	right StringNode
	op    BinaryOp
}

func (node *BinaryStringExprNode) Accept(visitor Visitor) {
	visitor.VisitBinaryStringExprNodeStart(node)
	node.left.Accept(visitor)
	node.right.Accept(visitor)
	visitor.VisitBinaryStringExprNodeEnd(node)
}

func (*BinaryStringExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *BinaryStringExprNode) EvalBool(s Symbols) (bool, error) {
	leftResult, err := node.left.EvalString(s)
	if err != nil {
		return false, err
	}
	rightResult, err := node.right.EvalString(s)
	if err != nil {
		return false, err
	}

	if leftResult == nil || rightResult == nil {
		return false, nil
	}

	switch node.op {
	case BinaryOpEQ:
		return *leftResult == *rightResult, nil
	case BinaryOpNEQ:
		return *leftResult != *rightResult, nil
	case BinaryOpLT:
		return *leftResult < *rightResult, nil
	case BinaryOpLTE:
		return *leftResult <= *rightResult, nil
	case BinaryOpGT:
		return *leftResult > *rightResult, nil
	case BinaryOpGTE:
		return *leftResult >= *rightResult, nil
	case BinaryOpContains:
		return strings.Contains(*leftResult, *rightResult), nil
	case BinaryOpNotContains:
		return !strings.Contains(*leftResult, *rightResult), nil
	}

	return false, errors.Errorf("unhandled string binary expression type %v", node.op)
}

func (node *BinaryStringExprNode) String() string {
	return fmt.Sprintf("%v %v %v", node.left, binaryOpNames[node.op], node.right)
}

type IsNilExprNode struct {
	symbol SymbolNode
	op     BinaryOp
}

func (node *IsNilExprNode) Accept(visitor Visitor) {
	visitor.VisitIsNilExprNodeStart(node)
	node.symbol.Accept(visitor)
	visitor.VisitIsNilExprNodeEnd(node)
}

func (*IsNilExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *IsNilExprNode) EvalBool(s Symbols) (bool, error) {
	isNil, err := s.IsNil(node.symbol.Symbol())
	if err != nil {
		return false, err
	}

	switch node.op {
	case BinaryOpEQ:
		return isNil, nil
	case BinaryOpNEQ:
		return !isNil, nil
	}

	return false, errors.Errorf("unhandled binary expression type %v", node)
}

func (node *IsNilExprNode) String() string {
	return fmt.Sprintf("%v %v null", node.symbol, binaryOpNames[node.op])
}

func NewInt64BetweenOp(nodes []Int64Node) (*Int64BetweenExprNode, error) {
	if len(nodes) != 3 {
		return nil, errors.Errorf("incorrect number of values provided to Int64BetweenExprNode: %v", len(nodes))
	}
	return &Int64BetweenExprNode{
		left:  nodes[0],
		lower: nodes[1],
		upper: nodes[2],
	}, nil
}

type Int64BetweenExprNode struct {
	left  Int64Node
	lower Int64Node
	upper Int64Node
}

func (node *Int64BetweenExprNode) Accept(visitor Visitor) {
	visitor.VisitInt64BetweenExprNodeStart(node)
	node.left.Accept(visitor)
	node.lower.Accept(visitor)
	node.upper.Accept(visitor)
	visitor.VisitInt64BetweenExprNodeEnd(node)
}

func (*Int64BetweenExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *Int64BetweenExprNode) EvalBool(s Symbols) (bool, error) {
	leftResult, err := node.left.EvalInt64(s)
	if err != nil {
		return false, err
	}

	if leftResult == nil {
		return false, err
	}

	lowerResult, err := node.lower.EvalInt64(s)
	if err != nil {
		return false, err
	}
	if lowerResult == nil {
		return false, nil
	}

	upperResult, err := node.upper.EvalInt64(s)
	if err != nil {
		return false, err
	}
	if upperResult == nil {
		return false, nil
	}

	return *leftResult >= *lowerResult && *leftResult < *upperResult, nil
}

func (node *Int64BetweenExprNode) String() string {
	return fmt.Sprintf("%v between %v and %v", node.left, node.lower, node.upper)
}

func NewFloat64BetweenOp(nodes []Float64Node) (*Float64BetweenExprNode, error) {
	if len(nodes) != 3 {
		return nil, errors.Errorf("incorrect number of values provided to Float64BetweenExprNode: %v", len(nodes))
	}
	return &Float64BetweenExprNode{
		left:  nodes[0],
		lower: nodes[1],
		upper: nodes[2],
	}, nil
}

type Float64BetweenExprNode struct {
	left  Float64Node
	lower Float64Node
	upper Float64Node
}

func (node *Float64BetweenExprNode) Accept(visitor Visitor) {
	visitor.VisitFloat64BetweenExprNodeStart(node)
	node.left.Accept(visitor)
	node.lower.Accept(visitor)
	node.upper.Accept(visitor)
	visitor.VisitFloat64BetweenExprNodeEnd(node)
}

func (*Float64BetweenExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *Float64BetweenExprNode) EvalBool(s Symbols) (bool, error) {
	leftResult, err := node.left.EvalFloat64(s)
	if err != nil {
		return false, err
	}

	if leftResult == nil {
		return false, err
	}

	lowerResult, err := node.lower.EvalFloat64(s)
	if err != nil {
		return false, err
	}
	if lowerResult == nil {
		return false, nil
	}

	upperResult, err := node.upper.EvalFloat64(s)
	if err != nil {
		return false, err
	}
	if upperResult == nil {
		return false, nil
	}

	return *leftResult >= *lowerResult && *leftResult < *upperResult, nil
}

func (node *Float64BetweenExprNode) String() string {
	return fmt.Sprintf("%v between %v and %v", node.left, node.lower, node.upper)
}

type DatetimeBetweenExprNode struct {
	left  DatetimeNode
	lower DatetimeNode
	upper DatetimeNode
}

func (node *DatetimeBetweenExprNode) Accept(visitor Visitor) {
	visitor.VisitDatetimeBetweenExprNodeStart(node)
	node.left.Accept(visitor)
	node.lower.Accept(visitor)
	node.upper.Accept(visitor)
	visitor.VisitDatetimeBetweenExprNodeEnd(node)
}

func (*DatetimeBetweenExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *DatetimeBetweenExprNode) EvalBool(s Symbols) (bool, error) {
	leftResult, err := node.left.EvalDatetime(s)
	if err != nil {
		return false, err
	}

	if leftResult == nil {
		return false, err
	}

	lowerResult, err := node.lower.EvalDatetime(s)
	if err != nil {
		return false, err
	}
	if lowerResult == nil {
		return false, nil
	}

	upperResult, err := node.upper.EvalDatetime(s)
	if err != nil {
		return false, err
	}
	if upperResult == nil {
		return false, nil
	}

	return (leftResult.Equal(*lowerResult) || leftResult.After(*lowerResult)) && leftResult.Before(*upperResult), nil
}

func (node *DatetimeBetweenExprNode) String() string {
	return fmt.Sprintf("%v between %v and %v", node.left, node.lower, node.upper)
}