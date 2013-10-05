package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	TYPE_CONS = iota
	TYPE_VALUE
)

type SExpression interface {
	ExprType() int
	ToString() string
}

type ConsCell struct {
	Car SExpression
	Cdr *ConsCell
}

func (cons *ConsCell) ExprType() int {
	return TYPE_CONS
}

func (cons *ConsCell) ToString() string {
	s := "("

	for ; cons != nil; cons = cons.Cdr {
		if cons.Car != nil {
			s += cons.Car.ToString()
		} else {
			s += "nil"
		}

		if cons.Cdr != nil {
			s += " "
		}
	}

	s += ")"
	return s
}

type Value struct {
	Val string
}

func (val *Value) ExprType() int {
	return TYPE_VALUE
}

func (val *Value) ToString() string {
	return val.Val
}

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

func parseValue(s string) (expr SExpression, rest string, err error) {
	i, size := 0, len(s)
	valBytes := []byte{}

	for ; i < size && s[i] != ')' && !isWhiteSpace(s[i]); i++ {
		valBytes = append(valBytes, s[i])
	}

	if len(valBytes) < 1 {
		return nil, s, errors.New("Value missing")
	}

	return &Value{string(valBytes)}, s[i:], nil
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

func readFile() (s string, err error) {
	var content []byte
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	content, err = ioutil.ReadFile(filename)
	if err == nil {
		s = string(content)
	}
	return
}

func main() {
	s, err := readFile()
	if err == nil {
		expr, rest, parseErr := parse(s)
		if parseErr != nil {
			fmt.Println(parseErr)
		} else {
			fmt.Println(expr.ToString())
			fmt.Println("----")
			fmt.Println(string(rest))
		}
	} else {
		fmt.Println("read error")
	}
}
