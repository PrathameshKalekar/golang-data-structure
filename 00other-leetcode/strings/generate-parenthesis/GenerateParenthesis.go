/*
22 Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]
*/
package main

import "fmt"

func main() {
	answer := generateParenthesis(3)
	fmt.Println(answer)
}

func generateParenthesis(n int) []string {
	var result []string
	var backtrack func(current string, open, close int)
	//for back tracking
	backtrack = func(current string, open, close int) {
		//if len is double the n then we git the valid parenthesis of len n
		if len(current) == 2*n {
			result = append(result, current)
			return
		}
		//ensure to add n starting parenthesis
		if open < n {
			backtrack(current+"(", open+1, close)
		}
		//ensures that for all n  starting parenthesis there is n closing parenthesis
		if close < open {
			backtrack(current+")", open, close+1)
		}
	}
	//backtrack from start
	backtrack("", 0, 0)
	//return the result
	return result
}
