package main

import (
	"fmt"
	"strconv"
	"strings"
)

func detectFunction(name, arg string) {
	switch name {
	case "print":
		fmt.Println(arg)
	}
}

func main() {
	s := `x=1.28
print(x)y=2.89
z=x+y
print(z)`

	mMemory := make(map[string]string)
	mFuncs := make(map[string]string)

	index := 0
	state := 0
	temp := ""

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '=' {
			temp = strings.TrimSpace(s[index:i])
			index = i
			state = 0
		} else if c == '(' {
			temp = strings.TrimSpace(s[index:i])
			index = i
		} else if c == ')' {
			value := s[index+1 : i]
			detectFunction(temp, mMemory[value])
			mFuncs[temp] = value
			state = 1
			index = i + 1
		} else if c == '+' {
			value := strings.TrimSpace(s[index+1 : i])
			mMemory[temp] = mMemory[value]
			index = i
			state = 2
		} else if c == '\n' {
			if index > 0 {
				value := strings.TrimSpace(s[index+1 : i])
				if state == 0 {
					mMemory[temp] = value
				} else if state == 2 {
					f, _ := strconv.ParseFloat(mMemory[temp], 32)
					f2, _ := strconv.ParseFloat(mMemory[value], 32)
					mMemory[temp] = fmt.Sprintf("%f", f+f2)
				}
			}
			index = i + 1
		}
	}
}
