package main

import (
	"fmt"
)

func main() {
	s := `print("It is    #$^@!~time?   !#9")`

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
}
