package main

import "fmt"

func main() {
	st := 'a'
	as := int(st)
	ss := rune(as)
	fmt.Printf("%d = %c", as, ss) // %c converts ascci to character

}
