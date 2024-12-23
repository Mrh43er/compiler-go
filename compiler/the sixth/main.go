package main

import (
	"fmt"
	"strings"
)

func main() {
	s := `          mnxn      =      "1  erherth  .42"`
	signs := "="
	list := []string{}
	mapping := map[string]string{}

	index := 0
	state := 0

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
					str := strings.TrimSpace(s[index:i])
					list = append(list, str)
				}

				if c != '"' && c != ' ' {
					list = append(list, string(c))
				}

				index = i + 1
			}
		}
	}
	str := strings.TrimSpace(s[index:])
	list = append(list, str)

	if list[1] == "=" {
		mapping[list[0]] = list[2]
	}

	fmt.Println(mapping)
}
