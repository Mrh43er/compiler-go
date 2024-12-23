package main

import (
	"fmt"
	"strings"
)

func main() {
	s := `if(x<y) then print   ("yes ()      ? kjjsw  $hbkj!")`
	signs := "()< \""

	state := 0
	index := 0

	for i := 0; i < len(s); i++ {
		c := s[i]
		if strings.ContainsRune(signs, rune(c)) {
			if c == '"' {
				if state == 0 {
					state = 1
					index++
				} else {
					state = 0
				}
			}
			if state == 0 {
				if i-index >= 1 {
					fmt.Println(s[index:i])
				}

				if c != '"' && c != ' ' {
					fmt.Println(string(c))
				}

				index = i + 1
			}
		}
	}
}
