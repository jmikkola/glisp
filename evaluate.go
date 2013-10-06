package main

import (
	"errors"
	"fmt"
)

func (val *Value) Evaluate() (SExpression, error) {
	return val, nil
}

func getFunctionName(head SExpression) (string, error) {
	if head == nil {
		return "", errors.New("null function name")
	}

	if value, ok := head.(*Value); ok {
		return value.Val, nil
	}

	return "", errors.New("non-symbol as function")
}

func (cons *ConsCell) Evaluate() (SExpression, error) {
	functionName, err := getFunctionName(cons.Car)
	if err != nil {
		return nil, err
	}

	builtins := map[string]func(args []SExpression) (SExpression, error){}
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
