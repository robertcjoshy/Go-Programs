package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Print("enter the number: ")
	fmt.Scan(&num)

	fmt.Printf("binary=%v\n", strconv.FormatInt(int64(num), 2))
	fmt.Printf("octal=%v\n", strconv.FormatInt(int64(num), 8))
	fmt.Printf("hexadecimal=%v", strconv.FormatInt(int64(num), 16))
}
