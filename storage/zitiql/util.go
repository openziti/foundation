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

package zitiql

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"regexp"
	"strings"
	"time"
)

func parse(str string, l ZitiQlListener, el antlr.ErrorListener, debug bool) {
	input := antlr.NewInputStream(str)
	lexer := NewZitiQlLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := NewZitiQlParser(stream)

	if debug {
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	} else {
		p.RemoveErrorListeners()
	}

	p.AddErrorListener(el)

	p.BuildParseTrees = true
	tree := p.Start()

	antlr.ParseTreeWalkerDefault.Walk(l, tree)
}

func Parse(str string, l ZitiQlListener) []ParseError {
	el := newErrorListener()
	parse(str, l, el, false)

	return el.Errors
}

func ParseWithDebug(str string, l ZitiQlListener) []ParseError {
	el := newErrorListener()
	parse(str, l, el, true)

	return el.Errors
}

type ParseError struct {
	Line    int
	Column  int
	Symbol  string
	Message string
}

func (p ParseError) Error() string {
	return fmt.Sprintf("%v. line: %v, column: %v, symbol: %v", p.Message, p.Line, p.Column, p.Symbol)
}

func newErrorListener() *ErrorListener {
	return &ErrorListener{
		Errors: []ParseError{},
	}
}

type ErrorListener struct {
	Errors []ParseError
}

func (el *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	s, ok := offendingSymbol.(*antlr.CommonToken)
	symbol := "<unknown>"
	if ok {
		symbol = s.GetText()
	}

	el.Errors = append(el.Errors, ParseError{
		Line:    line,
		Column:  column,
		Symbol:  symbol,
		Message: fmt.Sprintf(`Unexpected symbol: "%s" at line: %d column: %d`, s.GetText(), line, column),
	})
}

func (el *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// ignored
}

func (el *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// ignored
}

func (el *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	// ignored
}

func ParseZqlString(text string) string {
	t := strings.TrimSuffix(strings.TrimPrefix(text, `"`), `"`)

	//remove golang string back slash escaping
	t = strings.Replace(t, `\\`, `\`, -1)

	//remove ZitiQL string escaping
	t = strings.Replace(t, `\"`, `"`, -1)
	t = strings.Replace(t, `\f`, "\f", -1)
	t = strings.Replace(t, `\n`, "\n", -1)
	t = strings.Replace(t, `\r`, "\r", -1)
	t = strings.Replace(t, `\t`, "\t", -1)
	t = strings.Replace(t, `\\`, `\`, -1)

	return t
}

var dateTimeStripper = regexp.MustCompile(`^\s*datetime\(\s*(.*?)\s*\)\s*$`)

func ParseZqlDatetime(text string) (time.Time, error) {
	m := dateTimeStripper.FindAllStringSubmatch(text, -1)

	if m == nil || len(m) != 1 || len(m[0]) != 2 {
		return time.Time{}, fmt.Errorf("could not parse datetime (%s)", text)
	}

	//RFC3339 allows 'z','Z','t', and 'T'. Go's implementation only allows 'Z' and 'T'
	s := strings.Replace(m[0][1], "z", "Z", 1)
	s = strings.Replace(s, "t", "T", 1)

	t, err := time.Parse(time.RFC3339, s)

	return t, err
}

