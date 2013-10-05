package main

import (
	"errors"
)

func isWhiteSpace(b byte) bool {
	return b == ' ' || b == '\n' || b == '\t' || b == '\r'
}

func parseList(s string) (expr SExpression, rest string, err error) {
	var out *ConsCell = nil
	items := []SExpression{}

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
		out = &ConsCell{Car: items[i], Cdr: out}
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

func readQuotedString(s string) (result string, rest string, err error) {
	outChars := []rune{'"'}
	input := []rune(s)
	lastCh := ' '
	i, size := 1, len(input)

	if size < 2 || input[0] != '"' {
		return "", s, errors.New("Bad quoted string")
	}

	for ; i < size; i++ {
		ch := input[i]
		if lastCh == '\\' {
			outChars = append(outChars, escape(ch))
		} else if ch != '\\' {
			outChars = append(outChars, ch)
			if ch == '"' {
				break
			}
		}
		lastCh = ch
	}

	if input[i] != '"' {
		return "", s, errors.New("Bad quoted string")
	}

	return string(outChars), string(input[i+1:]), nil
}

func readSymbol(s string) (result string, rest string, err error) {
	i, size := 0, len(s)
	for ; i < size && s[i] != ')' && !isWhiteSpace(s[i]); i++ {
	}
	if i < 1 {
		return "", s, errors.New("Empty symbol")
	}
	return s[:i], s[i:], nil
}

func parseValue(s string) (expr SExpression, rest string, err error) {
	var result string
	if len(s) < 1 {
		return nil, s, errors.New("Missing value")
	}

	if s[0] == '"' {
		result, rest, err = readQuotedString(s)
	} else {
		result, rest, err = readSymbol(s)
	}

	if err == nil {
		expr = &Value{result}
	}

	return
}

func parse(s string) (expr SExpression, rest string, err error) {
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

	return parseValue(s[i:])
}

func ParseSExpression(s string) (expr SExpression, rest string, err error) {
    return parse(s)
}
