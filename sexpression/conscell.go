package sexpression

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