/*
1189  Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.

You can use each character in text at most once. Return the maximum number of instances that can be formed.



Example 1:



Input: text = "nlaebolko"
Output: 1
Example 2:



Input: text = "loonbalxballpoon"
Output: 2
Example 3:

Input: text = "leetcode"
Output: 0
*/

package main

import (
	"fmt"
)

func main() {
	string := "nlaeboko"
	answer := maximumNumbersOfBallons(string)
	fmt.Println(answer)
}

func maximumNumbersOfBallons(text string) int {
	balloonMap := make(map[rune]int)
	for _, value := range text {
		balloonMap[value]++
	}
	countB := balloonMap['b']
	countA := balloonMap['a']
	countL := balloonMap['l'] / 2
	countO := balloonMap['o'] / 2
	countN := balloonMap['n']
	min := countB
	if min > countA {
		min = countA
	}
	if min > countL {
		min = countL
	}
	if min > countO {
		min = countO
	}
	if min > countN {
		min = countN
	}
	return min
}
