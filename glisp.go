package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFile() (s string, err error) {
	var content []byte
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	content, err = ioutil.ReadFile(filename)
	if err == nil {
		s = string(content)
	}
	return
}

func main() {
	s, err := readFile()
	if err == nil {
		expr, _, parseErr := ParseSExpression(s)
		if parseErr != nil {
			fmt.Println(parseErr)
		} else {
			fmt.Println(expr.String())
			evaled, err := Evaluate(expr)
			if err == nil {
				fmt.Println("=>", evaled.String())
			} else {
				fmt.Println("Eval error", err)
			}
		}
	} else {
		fmt.Println("read error")
	}
}
