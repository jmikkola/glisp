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
}

type ConsCell struct {
	Car, Cdr *SExpression
}

func (cons *ConsCell) ExprType() int {
	return TYPE_CONS
}

type Value struct {
	Val string
}

func (val *Value) ExprType() int {
	return TYPE_VALUE
}

func parse(input chan rune) (SExpression, error) {
	chr, ok := <-input
	if !ok {
		return nil, errors.New("Unexpected end of input")
	}

	for isWhiteSpace(chr) {
		chr, ok = <-input
		if !ok {
			return nil, errors.New("Unexpected end of input")
		}
	}
	return nil, nil
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
	var sxp SExpression

	sxp = &ConsCell{}
	fmt.Println(sxp.ExprType())

	sxp = &Value{"val"}
	fmt.Println(sxp.ExprType())
}
