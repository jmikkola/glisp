package main

import (
	"errors"
	"fmt"

//	"unicode/utf8"
)

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
	outChars := []rune{}
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
		} else if ch == '"' {
			break
		} else if ch != '\\' {
			outChars = append(outChars, ch)
		}
		lastCh = ch
	}

	if input[i] != '"' {
		return "", s, errors.New("Bad quoted string")
	}

	return string(outChars), string(input[i+1:]), nil
}

func main() {
	result, rest, err := readQuotedString("\"a string \\n\\\"here!\" more stuff here")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
		fmt.Println(rest)
	}
}
