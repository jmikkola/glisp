package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var filename string

	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	content, err := ioutil.ReadFile(filename)
	if err == nil {
		s := string(content)
		fmt.Println(s)
	} else {
		fmt.Println(err)
	}
}
