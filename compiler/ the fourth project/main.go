package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "x=y+1.41/9+6"
	s := `print("It is true?!#")`
	// s := `if(x<y)then print("yes!")`

	state := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '"' {
			state = 1
			break
		}
	}

	if state == 1 {
		index := -1
		for i := 0; i < len(s); i++ {
			c := s[i]
			if c == '(' {
				fmt.Println(s[:i])
				fmt.Println("(")
			} else if c == '"' {
				if index >= 0 {
					fmt.Println(s[index+1 : i])
				}
				index = i
			} else if c == ')' {
				fmt.Println(")")
			}
		}
	} else {
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
		fmt.Println(strings.TrimSpace(s[index1:]))
	}
}
