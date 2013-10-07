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

func floatArgs(args []SExpression) ([]float64, error) {
	floatArgs := []float64{}

	for i, value := range args {
		switch value.ExprType() {
		case TYPE_INT:
			floatArgs = append(floatArgs, float64(value.(*Atom).Val.(GLInt)))
		case TYPE_FLOAT:
			floatArgs = append(floatArgs, float64(value.(*Atom).Val.(GLFloat)))
		default:
			return nil, errors.New(fmt.Sprintf("The %dth argument is not a number", i+1))
		}
	}

	return floatArgs, nil
}

var builtins map[string]func(args []SExpression) (SExpression, error) = map[string]func(args []SExpression) (SExpression, error){
	"+": func(args []SExpression) (SExpression, error) {
		floats, err := floatArgs(args)
		if err != nil {
			return nil, err
		}

		sum := 0.0
		for _, val := range floats {
			sum += val
		}

		return &Atom{TYPE_FLOAT, GLFloat(sum)}, nil
	},
	"-": func(args []SExpression) (SExpression, error) {
		floats, err := floatArgs(args)
		if err != nil {
			return nil, err
		}
		if len(floats) < 2 {
			return nil, errors.New("Cannot subtract with fewer than two numbers")
		}

		sum := floats[0]
		for _, val := range floats[1:] {
			sum -= val
		}

		return &Atom{TYPE_FLOAT, GLFloat(sum)}, nil
	},
	"*": func(args []SExpression) (SExpression, error) {
		floats, err := floatArgs(args)
		if err != nil {
			return nil, err
		}

		product := 1.0
		for _, val := range floats {
			product *= val
		}

		return &Atom{TYPE_FLOAT, GLFloat(product)}, nil
	},
	"/": func(args []SExpression) (SExpression, error) {
		floats, err := floatArgs(args)
		if err != nil {
			return nil, err
		}
		if len(floats) < 2 {
			return nil, errors.New("Cannot divide with fewer than two numbers")
		}

		numerator := floats[0]
		for _, val := range floats[1:] {
			numerator /= val
		}

		return &Atom{TYPE_FLOAT, GLFloat(numerator)}, nil
	},
	"list": func(args []SExpression) (SExpression, error) {
		var out *ConsCell = nil

		for i := len(args) - 1; i >= 0; i-- {
			out = &ConsCell{args[i], out}
		}

		return out, nil
	},
	"car": func(args []SExpression) (SExpression, error) {
		if len(args) != 1 {
			return nil, errors.New("car expects one argument")
		}

		cons, ok := args[0].(*ConsCell)
		if !ok {
			return nil, errors.New("car requires first argument to be a cons cell")
		}

		return cons.Car, nil
	},
	"cdr": func(args []SExpression) (SExpression, error) {
		if len(args) != 1 {
			return nil, errors.New("cdr expects one argument")
		}

		cons, ok := args[0].(*ConsCell)
		if !ok {
			return nil, errors.New("cdr requires first argument to be a cons cell")
		}

		return cons.Cdr, nil
	},
	"cons": func(args []SExpression) (SExpression, error) {
		if len(args) != 2 {
			return nil, errors.New("cons expects 2 arguments")
		}

		cons, ok := args[1].(*ConsCell)
		if !ok {
			return nil, errors.New("cons requires second argument to be a cons cell")
		}

		return &ConsCell{args[0], cons}, nil
	},
}

func (cons *ConsCell) Evaluate() (SExpression, error) {
	functionName, err := getFunctionName(cons.Car)
	if err != nil {
		return nil, err
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
