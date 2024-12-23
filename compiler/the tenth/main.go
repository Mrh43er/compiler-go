package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Global variables
var list [][]string
var memoryMap = make(map[string]int)

func applyOperator(left, right int, operator rune) int {
	switch operator {
	case '+':
		return left + right
	case '=':
		return right
	}
	return 0
}

func checkCondition(left, right int, operator string) bool {
	switch operator {
	case "<":
		return left < right
	case ">":
		return left > right
	}
	return false
}

func handleAssignment(tokens []string) {
	key := tokens[0]
	value := 0

	for i := 2; i < len(tokens); i += 2 {
		v, ok := memoryMap[tokens[i]]
		if !ok {
			v, _ = strconv.Atoi(tokens[i])
		}
		value = applyOperator(value, v, rune(tokens[i-1][0]))
	}
	memoryMap[key] = value
}

func handleFunction(tokens []string) {
	switch tokens[0] {
	case "print":
		val, ok := memoryMap[tokens[2]]
		if !ok {
			if v, err := strconv.Atoi(tokens[2]); err == nil {
				val = v
			} else {
				fmt.Println(tokens[2])
				return
			}
		}
		fmt.Println(val)
	}
}

func handleIfStatement(tokens []string) {
	// Extract left, right and operator
	left, ok := memoryMap[tokens[1]]
	if !ok {
		left, _ = strconv.Atoi(tokens[1])
	}

	operator := tokens[2]

	right, ok := memoryMap[tokens[3]]
	if !ok {
		right, _ = strconv.Atoi(tokens[3])
	}

	// Check condition
	if checkCondition(left, right, operator) {
		// If condition is true, execute the print statement
		fmt.Println("Condition met")
		handleFunction(tokens[4:])
	} else {
		fmt.Println("Condition not met")
	}
}

func handleForLoop(tokens []string) {
	iVar := tokens[1]
	from, ok := memoryMap[tokens[5]]
	if !ok {
		from, _ = strconv.Atoi(tokens[5])
	}
	memoryMap[iVar] = from

	to, ok := memoryMap[tokens[7]]
	if !ok {
		to, _ = strconv.Atoi(tokens[7])
	}

	// Print before the for loop starts
	fmt.Println("In for loop")
	for memoryMap[iVar] < to {
		// Print z in each iteration of the loop
		fmt.Println("Printing z:", memoryMap["z"])
		memoryMap[iVar]++
	}
}

func main() {
	// Open the input file
	file, err := os.Open("a.py")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Tokenize the file
	state := 0
	temp := ""
	list2 := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() + "\n"
		for _, c := range line {
			switch {
			case state == 1 && c == '"':
				list2 = append(list2, temp)
				temp = ""
				state = 0
			case state == 1:
				temp += string(c)
			case c == '=' || c == '+' || c == '(' || c == '<' || c == '>' || c == ')' || c == ',':
				if strings.TrimSpace(temp) != "" {
					list2 = append(list2, strings.TrimSpace(temp))
				}
				list2 = append(list2, string(c))
				temp = ""
			case c == ':':
				list2 = append(list2, ":")
				temp = ""
				state = 2
			case c == '"':
				state = 1
			case c == ' ':
				if strings.TrimSpace(temp) != "" {
					list2 = append(list2, strings.TrimSpace(temp))
					temp = ""
				}
			case c == '\n':
				if strings.TrimSpace(temp) != "" {
					list2 = append(list2, strings.TrimSpace(temp))
				}
				if state != 2 && len(list2) > 0 {
					list = append(list, append([]string{}, list2...))
					list2 = []string{}
				}
				temp = ""
				state = 0
			default:
				temp += string(c)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}

	// Process tokens
	for _, tokens := range list {
		if tokens[1] == "=" {
			handleAssignment(tokens)
		} else if tokens[0] == "if" {
			handleIfStatement(tokens)
		} else if tokens[0] == "for" {
			handleForLoop(tokens)
		}
	}

	// Optionally print memory map
	fmt.Println("Memory Map:", memoryMap)
}
