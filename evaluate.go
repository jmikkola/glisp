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

func (cons *ConsCell) Evaluate() (SExpression, error) {
	functionName, err := getFunctionName(cons.Car)
	if err != nil {
		return nil, err
	}

	builtins := map[string]func(args []SExpression) (SExpression, error){
		"+": func(args []SExpression) (SExpression, error) {
			//sum := 0.0

			return nil, nil
		},
	}
	_, ok := builtins[functionName]
	if ok {
		fmt.Println("found function ", functionName)
	}

	return nil, nil
}

func Evalute(sxp SExpression) (SExpression, error) {
	if sxp == nil {
		return nil, nil
	}

	return sxp.Evaluate()
}
