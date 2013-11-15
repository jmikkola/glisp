package sexpression

import (
	"strconv"
)

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
