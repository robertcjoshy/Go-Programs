package main

import "fmt"

func main() {
	var s rune = '7'
	if s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z' {
		fmt.Println("alphabet")
	} else {
		fmt.Println("not aplphabet")
	}
}
