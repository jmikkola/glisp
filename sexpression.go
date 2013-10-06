package main

const (
	TYPE_CONS = iota
	TYPE_VALUE
)

type SExpression interface {
	ExprType() int
	ToString() string
	Evaluate() (SExpression, error)
}

type ConsCell struct {
	Car SExpression
	Cdr *ConsCell
}

func (cons *ConsCell) ExprType() int {
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
	Val string
}

func (val *Atom) ExprType() int {
	return TYPE_VALUE
}

func (val *Atom) ToString() string {
	return val.Val
}
