package main

import "fmt"

func main() {
	var num int = 0
	var sum int = 0
	fmt.Print("enter the range to find the sum: ")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		sum += i
	}
	fmt.Println(sum)
}
