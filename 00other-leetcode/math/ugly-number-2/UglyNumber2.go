/*
264 An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

Given an integer n, return the nth ugly number.

Example 1:

Input: n = 10
Output: 12
Explanation: [1, 2, 3, 4, 5, 6, 8, 9, 10, 12] is the sequence of the first 10 ugly numbers.
Example 2:

Input: n = 1
Output: 1
Explanation: 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.
*/
package main

import "fmt"

func main() {
	answer := uglyNumber2(10)
	fmt.Println(answer)
}

func uglyNumber2(n int) int {
	if n <= 0 {
		return 0
	}

	uglyNumbers := make([]int, n)
	uglyNumbers[0] = 1

	i2, i3, i5 := 0, 0, 0
	next2, next3, next5 := 2, 3, 5

	for i := 1; i < n; i++ {
		nextUgly := min(next2, min(next3, next5))
		uglyNumbers[i] = nextUgly

		if nextUgly == next2 {
			i2++
			next2 = uglyNumbers[i2] * 2
		}
		if nextUgly == next3 {
			i3++
			next3 = uglyNumbers[i3] * 3
		}
		if nextUgly == next5 {
			i5++
			next5 = uglyNumbers[i5] * 5
		}
	}

	return uglyNumbers[n-1]
}
