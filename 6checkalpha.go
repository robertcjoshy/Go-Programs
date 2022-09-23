package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		var s []string // slice of string
		var c string // string

		fmt.Print("enter the char: ")
		fmt.Scan(&c)

		s = []string{"a", "e", "i", "o", "u"}
		fmt.Println(s, s[1])

		for _, k := range s {
			if k == c {
				fmt.Print("vowel")
				break
			} else {
				fmt.Print("consonant")
				break
			}
		}*/

	ss := "aeiou"
	var v string

	fmt.Print("enter the char: ")
	fmt.Scan(&v)

	s := strings.Contains(ss, v)

	if s {
		fmt.Print("vowel")
	} else {
		fmt.Print("consonant")
	}
	bb := strings.Split(ss, "")
	fmt.Println(bb)

}
