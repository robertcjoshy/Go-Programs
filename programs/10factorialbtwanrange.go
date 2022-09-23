package main

import (
	"fmt"
)

func main() {
	var start int
	var finish int

	fmt.Print("enter the starting range: ")
	fmt.Scan(&start)

	fmt.Print("enter the finishing range:  ")
	fmt.Scan(&finish)

	if start < 0 && finish <= 0 {
		err := fmt.Errorf("invalid range")
		fmt.Println(err.Error())
	}

}
