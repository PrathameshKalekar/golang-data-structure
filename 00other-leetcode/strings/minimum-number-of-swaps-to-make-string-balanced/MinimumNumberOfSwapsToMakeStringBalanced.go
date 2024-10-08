package main

import "fmt"

func main() {
	s := "][]["
	answer := minimumNumbersOfSwapsToMakeStringBalanced(s)
	fmt.Println(answer)
}
func minimumNumbersOfSwapsToMakeStringBalanced(s string) int {
	cnt := 0
	res := 0
	for _, r := range s {
		if r == '[' {
			cnt++
		} else {
			cnt--
		}
		if cnt < res {
			res = cnt
		}
	}
	return (-res + 1) / 2
}
