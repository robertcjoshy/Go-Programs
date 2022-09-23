package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int32
	var t = make(map[string]int)

	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(unsafe.Sizeof(t))

}
