package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
    filename := "input.gl"
	content, err := ioutil.ReadFile(filename)
	if err == nil {
        s := string(content)
		fmt.Println(s)
	} else {
		fmt.Println(err)
	}
}
