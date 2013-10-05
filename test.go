package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func asRunes(s string) ([]rune, error) {
	runes := make([]rune, 0)
	bytes := []byte(s)

	for i := 0; i < len(bytes); {
		r, size := utf8.DecodeRune(bytes[i:])
		if r == utf8.RuneError {
			return nil, errors.New("bad rune")
		}

		runes = append(runes, r)
		i += size
	}

	return runes, nil
}

func main() {
	s := "abc"
	fmt.Printf("%T\n", s[0])
	runes, err := asRunes(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(runes)
		fmt.Printf("%T\n", runes)
	}

	a := 'a'
	fmt.Printf("%T\n", a)
}
