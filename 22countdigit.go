package main

import "fmt"

func main() {
	var num int
	fmt.Print("enter the number: ")
	fmt.Scan(&num)

	count := 0
	for num != 0 {
		num = num / 10
		count++
	}
	fmt.Println(count)
}
