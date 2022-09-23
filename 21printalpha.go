package main

import "fmt"

func main() {
	var char rune = 'A'

	for i := char; i <= 'Z'; i++ {
		fmt.Printf("%s ", string(i))
	}
}
