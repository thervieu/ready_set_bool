package main

import (
	"fmt"
	"os"
	"strings"
)

func evaluateLogicalRPN(expression string, variables map[string]bool) (bool, error) {
	tokens := strings.Split(expression, "")

	stack := []bool{}

	for _, token := range tokens {
		switch token {
		case "&":
			if len(stack) < 2 {
				return false, fmt.Errorf("Insufficient operands for AND")
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, operand1 && operand2)
		case "|":
			if len(stack) < 2 {
				return false, fmt.Errorf("Insufficient operands for OR")
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, operand1 || operand2)
		case "!":
			if len(stack) < 1 {
				return false, fmt.Errorf("Insufficient operands for NOT")
			}
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, !operand1)
		case "^":
			if len(stack) < 2 {
				return false, fmt.Errorf("Insufficient operands for XOR")
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, operand1 != operand2)
		case ">":
			if len(stack) < 2 {
				return false, fmt.Errorf("Insufficient operands for IMPLIES")
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, !operand1 || operand2)
		case "=":
			if len(stack) < 2 {
				return false, fmt.Errorf("Insufficient operands for EQUIVALENCE")
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, operand1 == operand2)
		default:
			if val, ok := variables[token]; ok {
				stack = append(stack, val)
			} else {
				return false, fmt.Errorf("Invalid token: %s", token)
			}
		}
	}

	if len(stack) != 1 {
		return false, fmt.Errorf("Invalid RPN expression")
	}

	return stack[0], nil
}

func boolToInt(a bool) int {
	if a {
		return 1
	}
	return 0
}

func evaluateTruthTable(s string, variables []string) (bool, error) {
	numVars := len(variables)

	for _, v := range variables {
		fmt.Printf("| %s ", v)
	}
	fmt.Print("| = |\n")
	for i := 0; i <= len(variables); i++ {
		fmt.Print("|---")
	}
	fmt.Print("|\n")

	truthTable := make([]map[string]bool, 1<<uint(numVars))
	for i := 0; i < len(truthTable); i++ {
		truthTable[i] = make(map[string]bool)
		for j := 0; j < len(variables); j++ {
			truthTable[i][variables[j]] = (i>>j)&1 == 1
		}

		result, err := evaluateLogicalRPN(s, truthTable[i])
		if err != nil {
			return false, err
		}
		if result {
			return true, nil
		}

	}
	return false, nil
}

func isUpper(char rune) bool {
	return (char >= 'A' && char <= 'Z')
}

func contains(s string, slice []string) bool {
	for _, str := range slice {
		if str == s {
			return true
		}
	}
	return false
}

func extractVariables(expression string) []string {
	var variables []string

	for _, char := range expression {
		if isUpper(char) && !contains(string(char), variables) {
			variables = append(variables, string(char))
		}
	}
	return variables
}

func isSpecial(c rune) bool {
	return (c == '&' || c == '!' || c == '^' || c == '>' || c == '=')
}

func isValid(s string) bool {
	for _, char := range s {
		if !isUpper(char) && !isSpecial(char) {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: program string_to_evaluate")
		return
	}
	if !isValid(os.Args[1]) {
		fmt.Println("Usage: program string_to_evaluate")
	}

	expression := os.Args[1]
	variables := extractVariables(expression)

	truth, err := evaluateTruthTable(expression, variables)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%t\n", truth)
}
