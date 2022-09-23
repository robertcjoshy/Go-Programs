package main

import "fmt"

func main() {
	var a int = 0
	var b int = 1
	var c int = a + b
	var n int

	// user input -----------------
	fmt.Print("Ënter the number: ")
	fmt.Scan(&n)

	// print first and second number ----------------------
	fmt.Printf("%d\t%d\t", a, b)

	for c <= n {
		fmt.Printf("%d\t", c)

		// changing values of a and b ----------------
		a = b
		b = c
		c = a + b
	}
}
