package sexpression

import (
	"errors"
)

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