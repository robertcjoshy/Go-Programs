package main

import "fmt"

func main() {
	num := 0
	rev := 0
	fmt.Print("enter the number: ")
	fmt.Scan(&num)

	for num != 0 {
		temp := num % 10
		rev = rev*10 + temp
		num /= 10
	}
	fmt.Println(rev)
	fmt.Println(54321 % 10)
}
