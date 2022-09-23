package main

import "fmt"

func main() {
	num1 := 0
	num2 := 0
	factor := 0

	fmt.Print("enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("enter the first number: ")
	fmt.Scan(&num2)

	for i := 1; i <= num1 && i <= num2; i++ {
		if num1%i == 0 && num2%i == 0 {
			factor = i
		}
	}
	fmt.Println(factor)
}
