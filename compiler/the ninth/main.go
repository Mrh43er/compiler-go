package main

import (
	"fmt"
	"strings"
)

var list []string

func detectIf(s string) {
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
					list = append(list, s[index:i])
				}

				if c != '"' && c != ' ' {
					list = append(list, string(c))
				}

				index = i + 1
			}
		}
	}
}

func main() {
	s := ` x=1.28
y  =2.89
z=x+y
if(z  <y)   then print(    "       =  (dd ?<d ffff   ( ) yes  ? <")`

	index := 0
	state := 0
	var temp string

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '=' {
			temp = strings.TrimSpace(s[index:i])
			list = append(list, temp)
			list = append(list, "=")
			index = i
			state = 0
		} else if c == '+' {
			value := strings.TrimSpace(s[index+1 : i])
			list = append(list, value)
			list = append(list, "+")
			index = i
			state = 1
		} else if c == 'f' {
			if s[i-1] == 'i' {
				state = 2
				break
			}
		} else if c == '\n' {
			if index > 0 {
				value := strings.TrimSpace(s[index+1 : i])
				if state == 0 || state == 1 {
					list = append(list, value)
				} else if state == 2 {
					detectIf(strings.TrimSpace(s[index:]))
				}
			}
			index = i + 1
		}
	}

	if state == 2 {
		detectIf(strings.TrimSpace(s[index:]))
	}

	for i, item := range list {
		fmt.Printf("list[%d]: %s\n", i, item)
	}
}
