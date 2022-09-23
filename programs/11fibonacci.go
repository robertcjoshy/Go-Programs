package main

import (
	"fmt"
)

func main() {
	var a int = 0
	var b int = 1
	var n int

	// user input --------------
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)

	// printing first and second numbers
	fmt.Printf("%d\t%d\t", a, b)

	for i := 2; i < n; i++ {
		c := a + b
		a = b
		b = c

		// print fibonacci series
		fmt.Printf("%d\t", c)
	}
}
