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
)

func NewStringArrayNode(values []string) *StringArrayNode {
	result := &StringArrayNode{}
	for _, val := range values {
		result.values = append(result.values, &StringConstNode{value: val})
	}
	return result
}

// StringArrayNode encapsulates a string array
type StringArrayNode struct {
	values []StringNode
}

func (node *StringArrayNode) String() string {
	return fmt.Sprintf("%v", node.values)
}

func (*StringArrayNode) GetType() NodeType {
	return NodeTypeOther
}

func (node *StringArrayNode) Accept(visitor Visitor) {
	visitor.VisitStringArrayNodeStart(node)
	for _, child := range node.values {
		child.Accept(visitor)
	}
	visitor.VisitStringArrayNodeEnd(node)
}

func (node *StringArrayNode) AsStringArray() *StringArrayNode {
	return node
}

// Float64ArrayNode encapsulates a float64 array
type Float64ArrayNode struct {
	values []Float64Node
}

func (node *Float64ArrayNode) String() string {
	return fmt.Sprintf("%v", node.values)
}

func (node *Float64ArrayNode) GetType() NodeType {
	return NodeTypeOther
}

func (node *Float64ArrayNode) Accept(visitor Visitor) {
	visitor.VisitFloat64ArrayNodeStart(node)
	for _, child := range node.values {
		child.Accept(visitor)
	}
	visitor.VisitFloat64ArrayNodeEnd(node)
}

func (node *Float64ArrayNode) AsStringArray() *StringArrayNode {
	result := &StringArrayNode{}
	for _, child := range node.values {
		result.values = append(result.values, child)
	}
	return result
}

// Int64ArrayNode encapsulates an int64 array
type Int64ArrayNode struct {
	values []Int64Node
}

func (node *Int64ArrayNode) String() string {
	return fmt.Sprintf("%v", node.values)
}

func (node *Int64ArrayNode) GetType() NodeType {
	return NodeTypeOther
}

func (node *Int64ArrayNode) ToFloat64ArrayNode() *Float64ArrayNode {
	result := &Float64ArrayNode{}
	for _, intNode := range node.values {
		result.values = append(result.values, intNode.ToFloat64())
	}
	return result
}

func (node *Int64ArrayNode) Accept(visitor Visitor) {
	visitor.VisitInt64ArrayNodeStart(node)
	for _, child := range node.values {
		child.Accept(visitor)
	}
	visitor.VisitInt64ArrayNodeEnd(node)
}

func (node *Int64ArrayNode) AsStringArray() *StringArrayNode {
	result := &StringArrayNode{}
	for _, child := range node.values {
		result.values = append(result.values, child)
	}
	return result
}

// DatetimeArrayNode encapsulates a datetime array
type DatetimeArrayNode struct {
	values []DatetimeNode
}

func (node *DatetimeArrayNode) String() string {
	return fmt.Sprintf("%v", node.values)
}

func (node *DatetimeArrayNode) GetType() NodeType {
	return NodeTypeOther
}

func (node *DatetimeArrayNode) Accept(visitor Visitor) {
	visitor.VisitDatetimeArrayNodeStart(node)
	for _, child := range node.values {
		child.Accept(visitor)
	}
	visitor.VisitDatetimeArrayNodeEnd(node)
}

type InStringArrayExprNode struct {
	left  StringNode
	right *StringArrayNode
}

func (node *InStringArrayExprNode) String() string {
	return fmt.Sprintf("%v in %v", node.left, node.right)
}

func (*InStringArrayExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *InStringArrayExprNode) EvalBool(s Symbols) (bool, error) {
	left, err := node.left.EvalString(s)
	if err != nil {
		return false, err
	}

	for _, rightNode := range node.right.values {
		right, err := rightNode.EvalString(s)
		if err != nil {
			return false, err
		}
		if left != nil && right != nil {
			if *left == *right {
				return true, nil
			}
		}
	}
	return false, nil
}

func (node *InStringArrayExprNode) Accept(visitor Visitor) {
	visitor.VisitInStringArrayExprNodeStart(node)
	node.left.Accept(visitor)
	node.right.Accept(visitor)
	visitor.VisitInStringArrayExprNodeEnd(node)
}

type InInt64ArrayExprNode struct {
	left  Int64Node
	right *Int64ArrayNode
}

func (node *InInt64ArrayExprNode) String() string {
	return fmt.Sprintf("%v in %v", node.left, node.right)
}

func (node *InInt64ArrayExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *InInt64ArrayExprNode) EvalBool(s Symbols) (bool, error) {
	left, err := node.left.EvalInt64(s)
	if err != nil {
		return false, err
	}

	for _, rightNode := range node.right.values {
		right, err := rightNode.EvalInt64(s)
		if err != nil {
			return false, err
		}
		if left != nil && right != nil {
			if *left == *right {
				return true, nil
			}
		}
	}
	return false, nil
}

func (node *InInt64ArrayExprNode) Accept(visitor Visitor) {
	visitor.VisitInInt64ArrayExprNodeStart(node)
	node.left.Accept(visitor)
	node.right.Accept(visitor)
	visitor.VisitInInt64ArrayExprNodeEnd(node)
}

type InFloat64ArrayExprNode struct {
	left  Float64Node
	right *Float64ArrayNode
}

func (node *InFloat64ArrayExprNode) String() string {
	return fmt.Sprintf("%v in %v", node.left, node.right)
}

func (node *InFloat64ArrayExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *InFloat64ArrayExprNode) EvalBool(s Symbols) (bool, error) {
	left, err := node.left.EvalFloat64(s)
	if err != nil {
		return false, err
	}

	for _, rightNode := range node.right.values {
		right, err := rightNode.EvalFloat64(s)
		if err != nil {
			return false, err
		}
		if left != nil && right != nil {
			if *left == *right {
				return true, nil
			}
		}
	}
	return false, nil
}

func (node *InFloat64ArrayExprNode) Accept(visitor Visitor) {
	visitor.VisitInFloat64ArrayExprNodeStart(node)
	node.left.Accept(visitor)
	node.right.Accept(visitor)
	visitor.VisitInFloat64ArrayExprNodeEnd(node)
}

type InDatetimeArrayExprNode struct {
	left  DatetimeNode
	right *DatetimeArrayNode
}

func (node *InDatetimeArrayExprNode) String() string {
	return fmt.Sprintf("%v in %v", node.left, node.right)
}

func (*InDatetimeArrayExprNode) GetType() NodeType {
	return NodeTypeBool
}

func (node *InDatetimeArrayExprNode) EvalBool(s Symbols) (bool, error) {
	left, err := node.left.EvalDatetime(s)
	if err != nil {
		return false, err
	}

	for _, rightNode := range node.right.values {
		right, err := rightNode.EvalDatetime(s)
		if err != nil {
			return false, err
		}
		if left != nil && right != nil {
			if left.Equal(*right) {
				return true, nil
			}
		}
	}
	return false, nil
}

func (node *InDatetimeArrayExprNode) Accept(visitor Visitor) {
	visitor.VisitInDatetimeArrayExprNodeStart(node)
	node.left.Accept(visitor)
	node.right.Accept(visitor)
	visitor.VisitInDatetimeArrayExprNodeEnd(node)
}