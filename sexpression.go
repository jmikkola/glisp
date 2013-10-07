package main

import (
	"fmt"
	"strconv"
)

type GLType int

const (
	TYPE_CONS GLType = iota
	TYPE_INT
	TYPE_FLOAT
	TYPE_STRING
	TYPE_SYMBOL
)

type SExpression interface {
	ExprType() GLType
	ToString() string
	Evaluate() (SExpression, error)
}

type ConsCell struct {
	Car SExpression
	Cdr *ConsCell
}

func (cons *ConsCell) ExprType() GLType {
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

type Atom struct {
	Type GLType
	Val  fmt.Stringer
}

func (val *Atom) ExprType() GLType {
	return val.Type
}

func (atom *Atom) ToString() string {
	return atom.Val.String()
}

// Types for the values of an Atom:

type GLSymbol string
type GLString string
type GLFloat float64
type GLInt int64

func (sym GLSymbol) String() string {
	return string(sym)
}

func (str GLString) String() string {
	return "\"" + string(str) + "\""
}

func (gflt GLFloat) String() string {
	return strconv.FormatFloat(float64(gflt), 'g', -1, 64)
}

func (gint GLInt) String() string {
	return strconv.FormatInt(int64(gint), 10)
}
