package sexpression

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
