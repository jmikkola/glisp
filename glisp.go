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

func parseList(bytes []byte) (expr SExpression, rest []byte, err error) {
	//fmt.Println("parseList", string(bytes))
	var out *ConsCell = nil
	items := []SExpression{}

	for len(bytes) > 0 && bytes[0] != ')' {
		item, bx, err := parse(bytes)
		if err != nil {
			return nil, nil, err
		}
		items = append(items, item)
		bytes = bx
	}

	if len(bytes) < 1 {
		return nil, nil, errors.New("Unended list")
	}

	for i := len(items) - 1; i >= 0; i-- {
		out = &ConsCell{Car: items[i], Cdr: out}
	}

	return out, bytes, nil
}

func parseValue(bytes []byte) (expr SExpression, rest []byte, err error) {
	//fmt.Println("parseValue", string(bytes))
	i, size := 0, len(bytes)
	valBytes := []byte{}

	for ; i < size && bytes[i] != ')' && !isWhiteSpace(bytes[i]); i++ {
		valBytes = append(valBytes, bytes[i])
	}

	if len(valBytes) < 1 {
		return nil, nil, errors.New("Value missing")
	}

	//fmt.Println("parseValue: valBytes", string(valBytes))
	//fmt.Println("parseValue: rest", string(bytes[i:]))
	return &Value{string(valBytes)}, bytes[i+1:], nil
}

func parse(bytes []byte) (expr SExpression, rest []byte, err error) {
	//fmt.Println("parse", string(bytes))
	i, size := 0, len(bytes)
	for i < size && isWhiteSpace(bytes[i]) {
		i++
	}

	if i >= size {
		return nil, nil, errors.New("Unexpected end of input")
	}

	if bytes[i] == '(' {
		return parseList(bytes[i+1:])
	} else if bytes[i] == ')' {
		return nil, nil, errors.New("Unexpected end of list")
	}

	return parseValue(bytes[i:])
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
		bytes := []byte(s)
		expr, rest, parseErr := parse(bytes)
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
