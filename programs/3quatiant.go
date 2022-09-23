package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	fmt.Print("enter the dividend: ")
	fmt.Scan(&num1)
	fmt.Print("enter the diviser: ")
	fmt.Scan(&num2)
	result := num1 / num2
	result1 := num1 % num2
	fmt.Println(result)
	fmt.Println(result1)
}
