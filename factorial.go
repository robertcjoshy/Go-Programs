package main

import "fmt"

func main() {
	var n int
	f := 1

	// user input -----------
	fmt.Print("enter the number: ")
	fmt.Scan(&n)

	for i := n; i > 0; i-- {
		f = f * i
	}

	// print output-------------
	fmt.Println(f)
}
