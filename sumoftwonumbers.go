package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	// first input-----------------
	fmt.Print("enter the first number: ")
	fmt.Scan(&num1)

	// second input----------------
	fmt.Print("enter the second number: ")
	fmt.Scan(&num2)

	// adding two numbers--------
	sum := num1 + num2

	fmt.Println("sum of numbers =", sum)
}
