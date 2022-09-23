package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		//var num []int
		//var temp int
		var ss []int
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			fmt.Print(sc.Text())

		}
		fmt.Print()
	*/

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("start")
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		fmt.Println(input.Text())
		line := input.Text()
		fmt.Println(counts[line])
		counts[line] = counts[line] + 1
	}
	fmt.Println("---------------------------------------------")
	fmt.Println(counts)

	fmt.Println("-------------------------------------------")

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
