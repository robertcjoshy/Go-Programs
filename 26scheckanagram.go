package main

import "fmt"

func main() {
	var str1 string = "robert"
	var str2 string = "obrrte"
	flag := 0
	for i := 0; i <= len(str1)-1; i++ {
		if str1[i] != str2[i] {
			flag += 1
		} else {
			break
		}
	}
	fmt.Println(flag)
	if flag == len(str1) {
		fmt.Println("anagrams")
	}
}
