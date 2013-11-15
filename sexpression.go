package main

import (
	//	"fmt"
	"errors"
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

func GetTypeName(se SExpression) string {
	return TypeNames[se.ExprType()]
}

type Typer interface {
	AsFloat() (float64, error)
	AsInt() (int64, error)
	AsString() (string, error)
}

type SExpression interface {
	Typer
	ExprType() GLType
	String() string
	Evaluate() (SExpression, error)
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

func (a *BaseSExpression) AsFloat() (float64, error) {
	return 0.0, errors.New("Cannot convert to float: " + GetTypeName(a))
}

func (a *BaseSExpression) AsInt() (int64, error) {
	return 0, errors.New("Cannot convert to int: " + GetTypeName(a))
}

func (a *BaseSExpression) AsString() (string, error) {
	return "", errors.New("Cannot convert to string: " + GetTypeName(a))
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

// Symbol
type Symbol struct {
	Val string
	BaseSExpression
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
	BaseSExpression
}

func (_ *String) ExprType() GLType {
	return TYPE_STRING
}

func (s *String) String() string {
	return `"` + string(s.Val) + `"`
}

func (s *String) AsString() (string, error) {
	return s.Val, nil
}

// Float
type Float struct {
	Val float64
	BaseSExpression
}

func (_ *Float) ExprType() GLType {
	return TYPE_FLOAT
}

func (f *Float) String() string {
	return strconv.FormatFloat(f.Val, 'g', -1, 64)
}

func (f *Float) AsFloat() (float64, error) {
	return f.Val, nil
}

func (f *Float) AsInt() (int64, error) {
	return int64(f.Val), nil
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

func (i *Int) AsInt() (int64, error) {
	return i.Val, nil
}

func (i *Int) AsFloat() (float64, error) {
	return float64(i.Val), nil
}
