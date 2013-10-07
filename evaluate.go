package main

import (
	"errors"
	"fmt"

//	"strconv"
)

func (val *Atom) Evaluate() (SExpression, error) {
	return val, nil
}

func getFunctionName(head SExpression) (string, error) {
	if head == nil || head.ExprType() != TYPE_SYMBOL {
		return "", errors.New("non-symbol as function name")
	}

	atom, _ := head.(*Atom)
	return atom.Val.String(), nil
}

func expectType(expr SExpression, glType GLType) (err error) {
	if expr == nil || expr.ExprType() != glType {
		err = errors.New("Wrong type")
	}
	return
}

func argList(cons *ConsCell) ([]SExpression, error) {
	args := []SExpression{}

	for c := cons; c != nil; c = c.Cdr {
		value, err := Evaluate(c.Car)
		if err != nil {
			return nil, err
		} else {
			args = append(args, value)
		}
	}

	return args, nil
}

func (cons *ConsCell) Evaluate() (SExpression, error) {
	functionName, err := getFunctionName(cons.Car)
	if err != nil {
		return nil, err
	}

	builtins := map[string]func(args []SExpression) (SExpression, error){
		"+": func(args []SExpression) (SExpression, error) {
			// For now assume that it is all floats
			sum := 0.0

			for i, val := range args {
				// Also assuming that values are never nil (at least directly)
				switch val.ExprType() {
				case TYPE_INT:
					sum += float64(val.(*Atom).Val.(GLInt))
				case TYPE_FLOAT:
					sum += float64(val.(*Atom).Val.(GLFloat))
				default:
					return nil, errors.New(fmt.Sprintf("The %dth argument to + is not a number", i))
				}
			}

			return &Atom{TYPE_FLOAT, GLFloat(sum)}, nil
		},
	}

	fn, ok := builtins[functionName]
	if ok {
		args, err := argList(cons.Cdr)
		if err != nil {
			return nil, err
		}
		return fn(args)
	}

	return nil, errors.New("No function found called " + functionName)
}

func Evaluate(sxp SExpression) (SExpression, error) {
	if sxp == nil {
		return nil, nil
	}

	return sxp.Evaluate()
}
