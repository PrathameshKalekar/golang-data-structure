/*
150 You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:

The valid operators are '+', '-', '*', and '/'.
Each operand may be an integer or another expression.
The division between two integers always truncates toward zero.
There will not be any division by zero.
The input represents a valid arithmetic expression in a reverse polish notation.
The answer and all the intermediate calculations can be represented in a 32-bit integer.

Example 1:

Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
Example 2:

Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
Example 3:

Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	token := []string{"2", "1", "+", "3", "*"}
	answer := evaluateReversePolishNotation(token)
	fmt.Println(answer)
}

func evaluateReversePolishNotation(token []string) int {
	// Initialize an empty stack to store intermediate results
	stack := []int{}

	// Iterate over each token in the input array
	for _, value := range token {
		// Try to convert the token to an integer
		digit, err := strconv.Atoi(value)

		// If the conversion is successful, it's an operand
		if err != nil {
			// Pop the top two elements from the stack
			first, second := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			// Perform the operation based on the current operator token
			switch value {
			case "+":
				// Push the result of the addition onto the stack
				stack = append(stack, first+second)
			case "-":
				// Push the result of the subtraction onto the stack
				stack = append(stack, first-second)
			case "*":
				// Push the result of the multiplication onto the stack
				stack = append(stack, first*second)
			case "/":
				// Push the result of the division onto the stack
				stack = append(stack, first/second)
			}
			// Continue to the next token
			continue
		}

		// If it's an operand, push it onto the stack
		stack = append(stack, digit)
	}

	// The final result is the last remaining element on the stack
	return stack[0]
}
