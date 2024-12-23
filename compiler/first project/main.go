package main

import (
	"fmt"
)

func main() {
	// s := `S="Hello,I am fine!?yes."`
	s := `S="Hello,I am fine!?   yes.."`
	//String s = "S=\"Hello,I am fine!?yes.\"";

	var startString, endString int = -1, len(s)
	index1 := 0

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c == '"' {
			if startString >= 0 {
				endString = i
			} else {
				startString = i
				index1 = startString + 1
			}
		}

		if i > startString && i <= endString {
			if c == ' ' || c == ',' || c == '!' || c == '?' || c == '.' || i == endString {
				if i-index1 >= 1 {
					fmt.Println(s[index1:i])
				}
				index1 = i + 1
			}
		}
	}
}
