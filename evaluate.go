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

	atom, ok := head.(*Symbol)
	if !ok {
		return "", errors.New("non-symbol as function name")
	}

	return atom.String(), nil
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

func floatFunc(args []SExpression, minArgs int, fname string, fn func(args []float64) float64) (SExpression, error) {
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

	if len(floatArgs) < minArgs {
		return nil, errors.New(fmt.Sprintf("%s requires at least %d args, %d given", fname, minArgs, len(floatArgs)))
	}

	return &Float{Val: fn(floatArgs)}, nil
}

var builtins map[string]func(args []SExpression) (SExpression, error) = map[string]func(args []SExpression) (SExpression, error){
	"+": func(args []SExpression) (SExpression, error) {
		return floatFunc(args, 1, "+", func(args []float64) (sum float64) {
			for _, val := range args {
				sum += val
			}
			return
		})
	},
	"-": func(args []SExpression) (SExpression, error) {
		return floatFunc(args, 2, "-", func(args []float64) (sum float64) {
			sum = args[0]
			for _, val := range args[1:] {
				sum -= val
			}
			return
		})
	},
	"*": func(args []SExpression) (SExpression, error) {
		return floatFunc(args, 1, "*", func(args []float64) (product float64) {
			product = 1.0
			for _, val := range args {
				product *= val
			}
			return
		})
	},
	"/": func(args []SExpression) (SExpression, error) {
		return floatFunc(args, 2, "/", func(args []float64) (numerator float64) {
			numerator = args[0]
			for _, val := range args[1:] {
				numerator /= val
			}
			return
		})
	},
	"list": func(args []SExpression) (SExpression, error) {
		var out *ConsCell = nil

		for i := len(args) - 1; i >= 0; i-- {
			out = &ConsCell{Car: args[i], Cdr: out}
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

		return &ConsCell{Car: args[0], Cdr: cons}, nil
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
