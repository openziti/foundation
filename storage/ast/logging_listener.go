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
	zitiql "github.com/netfoundry/ziti-foundation/storage/zitiql"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"runtime"
	"strings"
)

type LoggingListener struct {
	PrintRuleLocation bool
	PrintChildren     bool
}

func (l *LoggingListener) printRuleLocationWithSkip(s int) {
	if l.PrintRuleLocation {
		pc, _, _, _ := runtime.Caller(s)
		f := runtime.FuncForPC(pc)
		s := strings.Split(f.Name(), ".")
		println(s[len(s)-1])
	}
}

func (l *LoggingListener) printChildren(tree antlr.ParseTree) {
	if l.PrintChildren {
		fmt.Printf("children for: %s\n", tree.GetText())

		for i, c := range tree.GetChildren() {
			fmt.Printf("-- %d: %s\n", i, c.(antlr.ParseTree).GetText())
		}
	}
}

func (l *LoggingListener) printDebug(tree antlr.ParseTree) {
	l.printRuleLocationWithSkip(2)
	l.printChildren(tree)
}

func (l LoggingListener) VisitTerminal(node antlr.TerminalNode) {
	l.printDebug(node)
}

func (l LoggingListener) VisitErrorNode(node antlr.ErrorNode) {
	l.printDebug(node)
}

func (l LoggingListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	l.printDebug(ctx)
}

func (l LoggingListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	l.printDebug(ctx)
}

func (l *LoggingListener) EnterQueryStmt(c *zitiql.QueryStmtContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterSortByExpr(c *zitiql.SortByExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterSortFieldExpr(c *zitiql.SortFieldExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitQueryStmt(c *zitiql.QueryStmtContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitSortByExpr(c *zitiql.SortByExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitSortFieldExpr(c *zitiql.SortFieldExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterSkipExpr(c *zitiql.SkipExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterLimitExpr(c *zitiql.LimitExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitSkipExpr(c *zitiql.SkipExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitLimitExpr(c *zitiql.LimitExprContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterSetFunction(c *zitiql.SetFunctionContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterBinary_lhs(c *zitiql.Binary_lhsContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterString_array(c *zitiql.String_arrayContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterNumber_array(c *zitiql.Number_arrayContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterDatetime_array(c *zitiql.Datetime_arrayContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterEnd(c *zitiql.EndContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterGroup(c *zitiql.GroupContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterOrConjunction(c *zitiql.OrConjunctionContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterOperationOp(c *zitiql.OperationOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterAndConjunction(c *zitiql.AndConjunctionContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterInStringArrayOp(c *zitiql.InStringArrayOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterInNumberArrayOp(c *zitiql.InNumberArrayOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterInDatetimeArrayOp(c *zitiql.InDatetimeArrayOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterBetweenNumberOp(c *zitiql.BetweenNumberOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterBetweenDateOp(c *zitiql.BetweenDateOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterBinaryLessThanNumberOp(c *zitiql.BinaryLessThanNumberOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterBinaryLessThanDatetimeOp(c *zitiql.BinaryLessThanDatetimeOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterBinaryGreaterThanNumberOp(c *zitiql.BinaryGreaterThanNumberOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterBinaryGreaterThanDatetimeOp(c *zitiql.BinaryGreaterThanDatetimeOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterBinaryEqualToStringOp(c *zitiql.BinaryEqualToStringOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterBinaryEqualToNumberOp(c *zitiql.BinaryEqualToNumberOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterBinaryEqualToDatetimeOp(c *zitiql.BinaryEqualToDatetimeOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterBinaryEqualToBoolOp(c *zitiql.BinaryEqualToBoolOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterBinaryEqualToNullOp(c *zitiql.BinaryEqualToNullOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) EnterBinaryContainsOp(c *zitiql.BinaryContainsOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitSetFunction(c *zitiql.SetFunctionContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitBinary_lhs(c *zitiql.Binary_lhsContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitString_array(c *zitiql.String_arrayContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitNumber_array(c *zitiql.Number_arrayContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitDatetime_array(c *zitiql.Datetime_arrayContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitEnd(c *zitiql.EndContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitGroup(c *zitiql.GroupContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitOrConjunction(c *zitiql.OrConjunctionContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitOperationOp(c *zitiql.OperationOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitAndConjunction(c *zitiql.AndConjunctionContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitInStringArrayOp(c *zitiql.InStringArrayOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitInNumberArrayOp(c *zitiql.InNumberArrayOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitInDatetimeArrayOp(c *zitiql.InDatetimeArrayOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitBetweenNumberOp(c *zitiql.BetweenNumberOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitBetweenDateOp(c *zitiql.BetweenDateOpContext) {
	l.printDebug(c)

}

func (l LoggingListener) ExitBinaryLessThanNumberOp(c *zitiql.BinaryLessThanNumberOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitBinaryLessThanDatetimeOp(c *zitiql.BinaryLessThanDatetimeOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitBinaryGreaterThanNumberOp(c *zitiql.BinaryGreaterThanNumberOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitBinaryGreaterThanDatetimeOp(c *zitiql.BinaryGreaterThanDatetimeOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitBinaryEqualToStringOp(c *zitiql.BinaryEqualToStringOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitBinaryEqualToNumberOp(c *zitiql.BinaryEqualToNumberOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitBinaryEqualToDatetimeOp(c *zitiql.BinaryEqualToDatetimeOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitBinaryEqualToBoolOp(c *zitiql.BinaryEqualToBoolOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitBinaryEqualToNullOp(c *zitiql.BinaryEqualToNullOpContext) {
	l.printDebug(c)
}

func (l LoggingListener) ExitBinaryContainsOp(c *zitiql.BinaryContainsOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBoolConst(c *zitiql.BoolConstContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBoolConst(c *zitiql.BoolConstContext) {
	l.printDebug(c)
}
