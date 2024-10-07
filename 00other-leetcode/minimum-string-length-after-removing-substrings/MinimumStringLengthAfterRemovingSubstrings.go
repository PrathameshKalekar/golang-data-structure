/*
2696 You are given a string s consisting only of uppercase English letters.

You can apply some operations to this string where, in one operation, you can remove any occurrence of one of the substrings "AB" or "CD" from s.

Return the minimum possible length of the resulting string that you can obtain.

Note that the string concatenates after removing the substring and could produce new "AB" or "CD" substrings.

Example 1:

Input: s = "ABFCACDB"
Output: 2
Explanation: We can do the following operations:
- Remove the substring "ABFCACDB", so s = "FCACDB".
- Remove the substring "FCACDB", so s = "FCAB".
- Remove the substring "FCAB", so s = "FC".
So the resulting length of the string is 2.
It can be shown that it is the minimum length that we can obtain.
Example 2:

Input: s = "ACBBD"
Output: 5
Explanation: We cannot do any operations on the string so the length remains the same.
*/
package main

import (
	"fmt"
)

func minimumStringLengthAfterRemovingSubstrings(s string) int {
	// Create a stack to store the characters
	stack := []rune{}

	// Traverse the string
	for _, char := range s {
		// If the stack has elements, check for the patterns "AB" or "CD"
		if len(stack) > 0 {
			top := stack[len(stack)-1] // Get the top of the stack

			// If top is 'A' and current char is 'B', pop the 'A'
			if top == 'A' && char == 'B' {
				stack = stack[:len(stack)-1] // Pop 'A'
				continue
			}

			// If top is 'C' and current char is 'D', pop the 'C'
			if top == 'C' && char == 'D' {
				stack = stack[:len(stack)-1] // Pop 'C'
				continue
			}
		}

		// Otherwise, push the current character onto the stack
		stack = append(stack, char)
	}

	// The length of the stack is the minimal length after removing substrings
	return len(stack)
}

func main() {
	// Test cases
	fmt.Println(minimumStringLengthAfterRemovingSubstrings("ABFCACDB")) // Output: 2
	fmt.Println(minimumStringLengthAfterRemovingSubstrings("ACBBD"))    // Output: 5
}
