package main

import (
	"errors"
	"fmt"
	"github.com/jmikkola/glisp/sexpression"
)

func getFunctionName(head sexpression.SExpression) (string, error) {
	if head == nil || head.ExprType() != TYPE_SYMBOL {
		return "", errors.New("non-symbol as function name")
	}

	atom, ok := head.(*Symbol)
	if !ok {
		return "", errors.New("non-symbol as function name")
	}

	return atom.String(), nil
}

func expectType(expr sexpression.SExpression, glType GLType) (err error) {
	if expr == nil || expr.ExprType() != glType {
		err = errors.New("Wrong type")
	}
	return
}

func argList(cons *sexpression.ConsCell) ([]sexpression.SExpression, error) {
	args := []sexpression.SExpression{}

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

func floatFunc(args []sexpression.SExpression, minArgs int, fname string, fn func(args []float64) float64) (sexpression.SExpression, error) {
	floatArgs := []float64{}

	for _, value := range args {
		floatVal, err := value.AsFloat()
		if err != nil {
			return nil, err
		}
		floatArgs = append(floatArgs, floatVal)
	}

	if len(floatArgs) < minArgs {
		return nil, errors.New(fmt.Sprintf("%s requires at least %d args, %d given", fname, minArgs, len(floatArgs)))
	}

	return &sexpression.Float{Val: fn(floatArgs)}, nil
}

var builtins = map[string]func(args []sexpression.SExpression) (sexpression.SExpression, error){
	"+": func(args []sexpression.SExpression) (sexpression.SExpression, error) {
		return floatFunc(args, 1, "+", func(args []float64) (sum float64) {
			for _, val := range args {
				sum += val
			}
			return
		})
	},
	"-": func(args []sexpression.SExpression) (sexpression.SExpression, error) {
		return floatFunc(args, 2, "-", func(args []float64) (sum float64) {
			sum = args[0]
			for _, val := range args[1:] {
				sum -= val
			}
			return
		})
	},
	"*": func(args []sexpression.SExpression) (sexpression.SExpression, error) {
		return floatFunc(args, 1, "*", func(args []float64) (product float64) {
			product = 1.0
			for _, val := range args {
				product *= val
			}
			return
		})
	},
	"/": func(args []sexpression.SExpression) (sexpression.SExpression, error) {
		return floatFunc(args, 2, "/", func(args []float64) (numerator float64) {
			numerator = args[0]
			for _, val := range args[1:] {
				numerator /= val
			}
			return
		})
	},
	"list": func(args []sexpression.SExpression) (sexpression.SExpression, error) {
		var out *sexpression.ConsCell = nil

		for i := len(args) - 1; i >= 0; i-- {
			out = &sexpression.ConsCell{Car: args[i], Cdr: out}
		}

		return out, nil
	},
	"car": func(args []sexpression.SExpression) (sexpression.SExpression, error) {
		if len(args) != 1 {
			return nil, errors.New("car expects one argument")
		}

		cons, ok := args[0].(*sexpression.ConsCell)
		if !ok {
			return nil, errors.New("car requires first argument to be a cons cell")
		}

		return cons.Car, nil
	},
	"cdr": func(args []sexpression.SExpression) (sexpression.SExpression, error) {
		if len(args) != 1 {
			return nil, errors.New("cdr expects one argument")
		}

		cons, ok := args[0].(*sexpression.ConsCell)
		if !ok {
			return nil, errors.New("cdr requires first argument to be a cons cell")
		}

		return cons.Cdr, nil
	},
	"cons": func(args []sexpression.SExpression) (sexpression.SExpression, error) {
		if len(args) != 2 {
			return nil, errors.New("cons expects 2 arguments")
		}

		cons, ok := args[1].(*sexpression.ConsCell)
		if !ok {
			return nil, errors.New("cons requires second argument to be a cons cell")
		}

		return &sexpression.ConsCell{Car: args[0], Cdr: cons}, nil
	},
}

func (cons *sexpression.ConsCell) Evaluate() (sexpression.SExpression, error) {
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

func Evaluate(sxp sexpression.SExpression) (sexpression.SExpression, error) {
	if sxp == nil {
		return nil, nil
	}

	return sxp.Evaluate()
}
