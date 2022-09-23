package main

import "fmt"

func main() {
	num1 := 0
	num2 := 0
	fmt.Print("enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("enter the first number: ")
	fmt.Scan(&num2)

	min := 0
	if num1 > num2 {
		min = num2
	} else {
		min = num1
	}

	factor := 0
	for true {
		if min%num1 == 0 && min%num2 == 0 {
			factor = min
			break
		}
		min = min + 1
	}
	fmt.Println(factor)
}
