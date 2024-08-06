/*
58 Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal substring consisting of non-space characters only.

Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
Example 2:

Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.
Example 3:

Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.
*/
package main

import "fmt"

func main() {
	s := ""
	answer := lengthOfLastWord(s)
	fmt.Println(answer)
}

func lengthOfLastWord(s string) int {
	//count to keep track of count
	count := 0
	//once the count is start mek it true to return the value once the word is ended
	start := false
	//range through the words reverses for last word
	for j := len(s) - 1; j >= 0; j-- {
		// ignore any spaces once the wrod count start make the start to true
		if s[j] != ' ' {
			start = true
			count++
		}
		//s[i] == " " and the count is started means the word is ended so return the count
		if s[j] == ' ' && start {
			return count
		}
	}
	//the there is only word in the string without space or no word present return the count
	return count
}
