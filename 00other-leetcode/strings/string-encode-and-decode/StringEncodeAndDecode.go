/*
271 Design an algorithm to encode a list of strings to a single string. The encoded string is then decoded back to the original list of strings.

# Please implement encode and decode

Example 1:

Input: ["neet","code","love","you"]

Output:["neet","code","love","you"]
Example 2:

Input: ["we","say",":","yes"]

Output: ["we","say",":","yes"]
*/
package main

import "fmt"

func main() {
	arr := []string{"neet", "code", "98", "you"}
	str := stringEncode(arr)
	answer := stringDecode(str)
	fmt.Println(answer)

}

func stringEncode(arr []string) string {
	encodedString := ""
	for _, value := range arr {
		encodedString += value + "`"

	}
	return encodedString
}

func stringDecode(s string) []string {
	start := 0
	result := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] == '`' {
			result = append(result, s[start:i])
			start = i + 1
		}
	}
	return result
}
