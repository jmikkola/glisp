package main

import (
	//	"fmt"
	"strconv"
)

type GLType int

const (
	TYPE_CONS GLType = iota
	TYPE_INT
	TYPE_FLOAT
	TYPE_STRING
	TYPE_SYMBOL
	TYPE_NIL
)

var TypeNames map[GLType]string = map[GLType]string{
	TYPE_CONS:   "cons",
	TYPE_INT:    "int",
	TYPE_FLOAT:  "float",
	TYPE_STRING: "string",
	TYPE_SYMBOL: "symbol",
	TYPE_NIL:    "nil",
}

type SExpression interface {
	ExprType() GLType
	String() string
	Evaluate() (SExpression, error)
}

func GetTypeName(se SExpression) string {
	return TypeNames[se.ExprType()]
}

// Provides "standard implementations" so that the types below always implement SExpression
type BaseSExpression struct{}

func (b *BaseSExpression) ExprType() GLType {
	panic("Not implemented")
}

func (b *BaseSExpression) String() string {
	panic("Not implemented")
}

func (b *BaseSExpression) Evaluate() (SExpression, error) {
	return b, nil
}

// The element that makes up lists
type ConsCell struct {
	Car SExpression
	Cdr *ConsCell
	BaseSExpression
}

func (cons *ConsCell) ExprType() GLType {
	return TYPE_CONS
}

func (cons *ConsCell) String() string {
	s := "("

	for ; cons != nil; cons = cons.Cdr {
		if cons.Car != nil {
			s += cons.Car.String()
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

// Atom will be used for default implementation of function to cast to specific types
type Atom struct {
	BaseSExpression
}

// Symbol
type Symbol struct {
	Val string
	Atom
}

func (_ *Symbol) ExprType() GLType {
	return TYPE_SYMBOL
}

func (s *Symbol) String() string {
	return s.Val
}

// String
type String struct {
	Val string
	Atom
}

func (_ *String) ExprType() GLType {
	return TYPE_STRING
}

func (s *String) String() string {
	return `"` + string(s.Val) + `"`
}

// Float
type Float struct {
	Val float64
	Atom
}

func (_ *Float) ExprType() GLType {
	return TYPE_FLOAT
}

func (f *Float) String() string {
	return strconv.FormatFloat(f.Val, 'g', -1, 64)
}

// Int
type Int struct {
	Val int64
	BaseSExpression
}

func (_ *Int) ExprType() GLType {
	return TYPE_INT
}

func (i *Int) String() string {
	return strconv.FormatInt(i.Val, 10)
}
