package main

import (
	"fmt"
	"os"
	"strings"
)

func evaluateLogicalRPN(expression string) (bool, error) {
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
			if token == "1" {
				stack = append(stack, true)
			} else if token == "0" {
				stack = append(stack, false)
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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./eval_formula string_to_evaluate")
		return
	}

	rpnExpression := os.Args[1]
	result, err := evaluateLogicalRPN(rpnExpression)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %t\n", result)
	}
}
