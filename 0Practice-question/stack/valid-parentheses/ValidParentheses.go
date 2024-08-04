/*
20 Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
*/
package main

import "fmt"

func main() {
	s := "(){}["
	answer := validParentheses(s)
	fmt.Println(answer)
}

func validParentheses(s string) bool {
	// Initialize a stack to keep track of opening brackets.
	stack := []rune{}

	// Create a map to match closing brackets with their corresponding opening brackets.
	matching := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Iterate over each character in the string.
	for _, char := range s {
		// If the character is an opening bracket, push it onto the stack.
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			// If the character is a closing bracket:
			// 1. Check if the stack is empty. If it is, the closing bracket has no matching opening bracket.
			// 2. Check if the top of the stack matches the corresponding opening bracket for the current closing bracket.
			if len(stack) == 0 || stack[len(stack)-1] != matching[char] {
				return false
			}
			// If it matches, pop the top of the stack.
			stack = stack[:len(stack)-1]
		}
	}

	// At the end, if the stack is empty, all brackets were matched correctly.
	// If not, some opening brackets didn't have matching closing brackets.
	return len(stack) == 0
}
