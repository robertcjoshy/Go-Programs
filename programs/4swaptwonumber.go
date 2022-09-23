package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	fmt.Print("enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("enter the second number: ")
	fmt.Scan(&num2)

	fmt.Println("first = ", num1)
	fmt.Println("second = ", num2)

	// swaping numbers---------------

	var temp int
	temp = num1
	num1 = num2
	num2 = temp

	// print numbers----------------

	fmt.Println(num1, num2)
}
