/*
65 Given a string s, return whether s is a valid number.

For example, all the following are valid numbers: "2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789", while the following are not valid numbers: "abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53".

Formally, a valid number is defined using one of the following definitions:

An integer number followed by an optional exponent.
A decimal number followed by an optional exponent.
An integer number is defined with an optional sign '-' or '+' followed by digits.

A decimal number is defined with an optional sign '-' or '+' followed by one of the following definitions:

Digits followed by a dot '.'.
Digits followed by a dot '.' followed by digits.
A dot '.' followed by digits.
An exponent is defined with an exponent notation 'e' or 'E' followed by an integer number.

The digits are defined as one or more digits.

Example 1:

Input: s = "0"

Output: true

Example 2:

Input: s = "e"

Output: false

Example 3:

Input: s = "."

Output: false
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Test cases
	tests := []string{
		"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3",
		"3e+7", "+6e-1", "53.5e93", "-123.456e789",
		"abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53",
	}

	for _, test := range tests {
		fmt.Printf("%q: %v\n", test, validNumber(test))
	}
}

func validNumber(s string) bool {

	// Trim any white spaces around the string
	s = strings.TrimSpace(s)

	// Flags for checking the structure
	dotSeen := false       // If a dot has been encountered
	eSeen := false         // If an exponent 'e' or 'E' has been encountered
	numberBeforeE := false // If digits have appeared before 'e' or 'E'
	numberAfterE := false  // If digits have appeared after 'e' or 'E'

	for i := 0; i < len(s); i++ {
		// Get the current character
		ch := s[i]

		// Handle digits
		if ch >= '0' && ch <= '9' {
			numberBeforeE = true // Digits are found before 'e'
			numberAfterE = true  // Also valid after 'e'

			// Handle dot '.'
		} else if ch == '.' {
			// If a dot or 'e' was already seen, return false
			if dotSeen || eSeen {
				return false
			}
			dotSeen = true

			// Handle exponent 'e' or 'E'
		} else if ch == 'e' || ch == 'E' {
			// If 'e' was already seen or no digits before 'e', return false
			if eSeen || !numberBeforeE {
				return false
			}
			eSeen = true
			numberAfterE = false // Expect numbers after 'e'

			// Handle sign '+' or '-'
		} else if ch == '+' || ch == '-' {
			// Sign can only be at the beginning or right after 'e'
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}

			// Any other character is invalid
		} else {
			return false
		}
	}

	// The number must have at least one digit before 'e' and valid digits after 'e'
	return numberBeforeE && (numberAfterE || !eSeen)
}
