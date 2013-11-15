package sexpression

import (
	"strconv"
)

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
