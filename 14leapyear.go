package main

import "fmt"

func main() {
	var year int = 1999
	if year%100 == 0 {
		if year%400 == 0 {
			fmt.Println("leapyear")
		} else {
			fmt.Println("not leap year")
		}
	} else if year%4 == 0 {
		fmt.Println("leapyear")
	} else {
		fmt.Println("not leap year")
	}
}
