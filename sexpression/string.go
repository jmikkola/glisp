package sexpression

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