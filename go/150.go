package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			}
			continue
		}
		n, _ := strconv.Atoi(token)
		stack = append(stack, n)
	}

	return stack[0]
}
