package main

import (
	"fmt"
)

const (
	sunday = iota << 1
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

func main() {

	fmt.Println(thursday)
}
