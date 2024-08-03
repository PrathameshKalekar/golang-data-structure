/*
Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
Example 2:

Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
*/
package main

import "fmt"

func main() {
	string := []string{"h", "e", "l", "l", "o"}
	fmt.Println("string before reverse : ", string)
	reverseString(string)
	fmt.Println("string after reverse : ", string)
}
func reverseString(s []string) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
