package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Println("start")
	fmt.Print("enter the number: ")
	fmt.Scan(&s)
	result := comma(s)
	fmt.Println(result)
}

func comma(s string) string {
	b := bytes.Buffer{}
	front := 0
	if s[front] == '+' || s[front] == '-' {
		front = 1
	}
	back := strings.Index(s, ".")
	if back == -1 {
		back = len(s)
	}
	final := s[front:back]
	pre := len(final) % 3
	if pre > 0 {
		b.Write([]byte(s[:pre]))
		if pre < len(final) {
			b.WriteString(",")
		}
	}
	for i, k := range final {
		if i%3 == 0 && i != 0 {
			b.WriteString(",")
		}
		b.WriteRune(k)
	}
	b.WriteString(s[back:])
	return b.String()
}
