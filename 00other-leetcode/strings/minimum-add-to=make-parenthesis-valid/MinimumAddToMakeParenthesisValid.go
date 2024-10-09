package main

import "fmt"

func main() {
	s := "())"
	answer := minimumAddToMakeParenthesisValid(s)
	fmt.Println(answer)
}

func minimumAddToMakeParenthesisValid(s string) int {
	moves := 0
	open := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if open < 0 {
				moves += open * -1
				open = 1
			} else {
				open++
			}
		} else {
			open--
		}
	}
	if open < 0 {
		open = open * -1
	}
	return moves + open
}
