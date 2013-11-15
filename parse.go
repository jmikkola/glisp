package main

import (
	"errors"
	"github.com/jmikkola/glisp/sexpression"
	"regexp"
	"strconv"
)

var intRe *regexp.Regexp = regexp.MustCompile("^-?\\d+$")
var floatRe *regexp.Regexp = regexp.MustCompile("^-?\\d+(\\.\\d+)?([eE][+-]?\\d+)?$")

func isWhiteSpace(b byte) bool {
	return b == ' ' || b == '\n' || b == '\t' || b == '\r'
}

func parseList(s string) (expr sexpression.SExpression, rest string, err error) {
	var out *sexpression.ConsCell = nil
	items := []sexpression.SExpression{}

	for len(s) > 0 && s[0] != ')' {
		item, rst, err := parse(s)
		if err != nil {
			return nil, s, err
		}
		items = append(items, item)
		s = rst
	}

	if len(s) < 1 {
		return nil, s, errors.New("Unended list")
	}

	for i := len(items) - 1; i >= 0; i-- {
		out = &sexpression.ConsCell{Car: items[i], Cdr: out}
	}

	return out, s[1:], nil
}

func escape(ch rune) rune {
	switch ch {
	case '"':
		return '"'
	case 'n':
		return '\n'
	case '\\':
		return '\\'
	}
	return ch
}

func readQuotedString(s string) (result sexpression.SExpression, rest string, err error) {
	outChars := []rune{}
	input := []rune(s)
	lastCh := ' '
	i, size := 1, len(input)

	if size < 2 || input[0] != '"' {
		return nil, s, errors.New("Bad quoted string")
	}

	for ; i < size; i++ {
		ch := input[i]
		if lastCh == '\\' {
			outChars = append(outChars, escape(ch))
		} else if ch == '"' {
			break
		} else if ch != '\\' {
			outChars = append(outChars, ch)
		}
		lastCh = ch
	}

	if input[i] != '"' {
		return nil, s, errors.New("Bad quoted string")
	}

	return &sexpression.String{Val: string(outChars)}, string(input[i+1:]), nil
}

func readWord(s string) (word string, rest string, err error) {
	i, size := 0, len(s)
	for ; i < size && s[i] != ')' && !isWhiteSpace(s[i]); i++ {
	}
	if i < 1 {
		return "", s, errors.New("Expected to read word")
	}
	return s[:i], s[i:], nil
}

func parseAtom(s string) (expr sexpression.SExpression, rest string, err error) {
	rest = s
	if len(s) < 1 {
		err = errors.New("Missing value")
		return
	}

	if s[0] == '"' {
		return readQuotedString(s)
	}

	word, rest, err := readWord(s)
	if err != nil {
		return
	}
	wordBytes := []byte(word)

	if intRe.Match(wordBytes) {
		intVal, err := strconv.ParseInt(word, 10, 64)
		if err == nil {
			expr = &sexpression.Int{Val: intVal}
		}
	} else if floatRe.Match(wordBytes) {
		floatVal, err := strconv.ParseFloat(word, 64)
		if err == nil {
			expr = &sexpression.Float{Val: floatVal}
		}
	} else {
		expr = &sexpression.Symbol{Val: word}
	}

	return
}

func parse(s string) (expr sexpression.SExpression, rest string, err error) {
	i, size := 0, len(s)
	for i < size && isWhiteSpace(s[i]) {
		i++
	}

	if i >= size {
		return nil, s, errors.New("Unexpected end of input")
	}

	if s[i] == '(' {
		return parseList(s[i+1:])
	} else if s[i] == ')' {
		return nil, s, errors.New("Unexpected end of list")
	}

	return parseAtom(s[i:])
}

func ParseSExpression(s string) (expr sexpression.SExpression, rest string, err error) {
	return parse(s)
}
