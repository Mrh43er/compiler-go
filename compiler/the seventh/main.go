package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "x = 1.289"
	s2 := "print(x)"
	list := []string{}
	signs := "="
	mMemory := make(map[string]string)
	mFuncs := make(map[string]string)

	index := 0
	state := 0

	for i := 0; i < len(s); i++ {
		c := s[i]
		if strings.ContainsRune(signs, rune(c)) {
			if state == 0 {
				if i-index >= 1 {
					str := strings.TrimSpace(s[index:i])
					list = append(list, str)
				}
				if c != ' ' {
					list = append(list, string(c))
				}
				index = i + 1
			}
		}
	}
	str := strings.TrimSpace(s[index:])
	list = append(list, str)

	if list[1] == "=" {
		mMemory[list[0]] = list[2]
	}

	index = -1
	for i := 0; i < len(s2); i++ {
		c := s2[i]
		if c == '(' {
			index = i + 1
		} else if c == ')' {
			if index >= 0 {
				mFuncs[s2[:index-1]] = s2[index:i]
			}
		}
	}

	if val, exists := mFuncs["print"]; exists {
		if memoryValue, exists := mMemory[val]; exists {
			fmt.Println(memoryValue)
		}
	}
}
