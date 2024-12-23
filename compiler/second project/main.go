package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "x=1.14"
	// s := "x=2+3"
	// s := "y="
	s := "x=1.423+ 2.46*9/8   /"

	index1 := 0

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '=' || c == '+' || c == '*' || c == '/' {
			if i-index1 >= 1 {
				fmt.Println(strings.TrimSpace(s[index1:i]))
			}
			fmt.Println(string(c))
			index1 = i + 1
		}
	}
	if index1 < len(s) {
		fmt.Println(strings.TrimSpace(s[index1:]))
	}
}
