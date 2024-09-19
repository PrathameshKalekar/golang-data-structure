/*
241 Given a string expression of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators. You may return the answer in any order.

The test cases are generated such that the output values fit in a 32-bit integer and the number of different results does not exceed 104.

Example 1:

Input: expression = "2-1-1"
Output: [0,2]
Explanation:
((2-1)-1) = 0
(2-(1-1)) = 2
Example 2:

Input: expression = "2*3-4*5"
Output: [-34,-14,-10,-10,10]
Explanation:
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10
*/
package main

import (
	"fmt"
	"strconv"
)

// Function to compute all possible results
func diffWaysToAddParentheses(expression string) []int {
	// Base case: if the expression is a number, return it as a single result
	if num, err := strconv.Atoi(expression); err == nil {
		return []int{num}
	}

	var results []int

	// Iterate through the string, and for every operator, split into left and right parts
	for i := 0; i < len(expression); i++ {
		char := expression[i]

		// If the character is an operator, split the expression
		if char == '+' || char == '-' || char == '*' {
			// Recursively solve left and right sub-expressions
			leftResults := diffWaysToAddParentheses(expression[:i])
			rightResults := diffWaysToAddParentheses(expression[i+1:])

			// Compute all combinations of results from left and right sub-expressions
			for _, left := range leftResults {
				for _, right := range rightResults {
					var res int
					switch char {
					case '+':
						res = left + right
					case '-':
						res = left - right
					case '*':
						res = left * right
					}
					results = append(results, res)
				}
			}
		}
	}

	return results
}

func main() {
	// Example 1
	expression1 := "2-1-1"
	fmt.Println("Input:", expression1)
	fmt.Println("Output:", diffWaysToAddParentheses(expression1))

	// Example 2
	expression2 := "2*3-4*5"
	fmt.Println("Input:", expression2)
	fmt.Println("Output:", diffWaysToAddParentheses(expression2))
}
