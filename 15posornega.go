package main

import "fmt"

func main() {
	var s int = -23
	if s == 0 {
		fmt.Println("please give valid number: ")
	} else if s > 0 {
		fmt.Println("positive")
	} else {
		fmt.Println("negative")
	}
}
