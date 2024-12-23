package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var memoryMap = make(map[string]interface{})
var functionMap = make(map[string][]string)
var lines [][]string

func handleFunction(list []string) {
	switch list[0] {
	case "print":
		value, exists := memoryMap[list[2]]
		if !exists {
			value = list[2]
		}
		fmt.Println(value)
	default:
		subLines, exists := functionMap[list[0]]
		if exists {
			memoryMap[subLines[1]] = list[2]
			handleLine(subLines[4:])
			delete(memoryMap, subLines[1])
		}
	}
}

func handleLine(line []string) {
	if line[0] == "def" {
		functionMap[line[1]] = line[2:]
	} else if len(line) > 2 && line[1] == "(" {
		handleFunction(line)
	}
}

func main() {
	file, err := os.Open("./d.py")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var list2 []string

	for scanner.Scan() {
		state := 0
		temp := ""
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			c := line[i]
			switch state {
			case 1:
				if c == '"' {
					list2 = append(list2, temp)
					temp = ""
					state = 0
				} else {
					temp += string(c)
				}
			case 0:
				switch c {
				case ' ', '=', '+', '(':
					if c == ' ' {
						if temp == "def" {
							list2 = append(list2, strings.TrimSpace(temp))
							temp = ""
						}
						continue
					}
					list2 = append(list2, strings.TrimSpace(temp))
					list2 = append(list2, string(c))
					temp = ""
				case ':':
					list2 = append(list2, ":")
					temp = ""
					state = 2
				case '<', '>', ')', ',':
					if strings.TrimSpace(temp) != "" {
						list2 = append(list2, strings.TrimSpace(temp))
					}
					list2 = append(list2, string(c))
					temp = ""
				case '"':
					state = 1
				default:
					temp += string(c)
				}
			}
		}
		if strings.TrimSpace(temp) != "" {
			list2 = append(list2, strings.TrimSpace(temp))
		}

		if state != 2 {
			if len(list2) > 0 {
				lines = append(lines, append([]string{}, list2...))
			}
			list2 = nil
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	for _, line := range lines {
		handleLine(line)
	}

	fmt.Println("Memory Map:", memoryMap)
	fmt.Println("Function Map:", functionMap)
}
