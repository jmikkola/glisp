package main

import (
	"errors"
	"fmt"
	"strconv"
)

func (val *Atom) Evaluate() (SExpression, error) {
	return val, nil
}

func getFunctionName(head SExpression) (string, error) {
	if head == nil {
		return "", errors.New("null function name")
	}

	if value, ok := head.(*Atom); ok {
		return value.Val, nil
	}

	return "", errors.New("non-symbol as function")
}

func asFloating(se SExpression) (float64, error) {
	if se == nil {
		return 0, errors.New("nil is not a number")
	}

	val, isVal := se.(*Atom)
	if !isVal {
		return 0, errors.New("numbers must be values")
	}

    return strconv.ParseFloat(val.Val, 64)
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
