/*
3043 You are given two arrays with positive integers arr1 and arr2.

A prefix of a positive integer is an integer formed by one or more of its digits, starting from its leftmost digit. For example, 123 is a prefix of the integer 12345, while 234 is not.

A common prefix of two integers a and b is an integer c, such that c is a prefix of both a and b. For example, 5655359 and 56554 have a common prefix 565 while 1223 and 43456 do not have a common prefix.

You need to find the length of the longest common prefix between all pairs of integers (x, y) such that x belongs to arr1 and y belongs to arr2.

Return the length of the longest common prefix among all pairs. If no common prefix exists among them, return 0.

Example 1:

Input: arr1 = [1,10,100], arr2 = [1000]
Output: 3
Explanation: There are 3 pairs (arr1[i], arr2[j]):
- The longest common prefix of (1, 1000) is 1.
- The longest common prefix of (10, 1000) is 10.
- The longest common prefix of (100, 1000) is 100.
The longest common prefix is 100 with a length of 3.
Example 2:

Input: arr1 = [1,2,3], arr2 = [4,4,4]
Output: 0
Explanation: There exists no common prefix for any pair (arr1[i], arr2[j]), hence we return 0.
Note that common prefixes between elements of the same array do not count.
*/
package main

import (
	"fmt"
	"strconv"
)

// Function to find the length of the longest common prefix between two strings
func commonPrefixLength(s1, s2 string) int {
	minLength := len(s1)
	if len(s2) < minLength {
		minLength = len(s2)
	}

	for i := 0; i < minLength; i++ {
		if s1[i] != s2[i] {
			return i
		}
	}
	return minLength
}

// Optimized function to find the longest common prefix between all pairs from arr1 and arr2
func longestCommonPrefix(arr1, arr2 []int) int {
	maxPrefixLength := 0

	// Early exit optimization: compare only numbers that start with the same digit
	for _, num1 := range arr1 {
		str1 := strconv.Itoa(num1) // Convert number to string
		for _, num2 := range arr2 {
			str2 := strconv.Itoa(num2) // Convert number to string

			// If the first digit doesn't match, skip this pair
			if str1[0] != str2[0] {
				continue
			}

			// Find the common prefix length between str1 and str2
			commonLength := commonPrefixLength(str1, str2)

			// Update the maximum prefix length found
			if commonLength > maxPrefixLength {
				maxPrefixLength = commonLength
			}
		}
	}

	return maxPrefixLength
}

func main() {
	// Example input (large input similar to what you might have faced TLE on)
	arr1 := []int{4464054, 6110338, 7722639, 1512348, 4467733, 4171071, 1431687, 7140215, 1967014, 2677727, 5617754, 8772784, 3525082, 4762316, 8205136, 4656162, 2348307, 4544604, 5948661, 7347242, 1464576, 7125228, 5710350, 9953233, 7629680, 8598961, 4324165, 6665234, 6476704, 6974090, 7550477, 9336784, 7711155, 2991977, 3613647, 8228591, 9124420, 1158251, 3244394, 8527011, 9159756, 5483480, 3220372, 1155125, 3350893, 6847308, 4513123, 6879731, 5915661, 4970866, 6899155, 9144876, 7360137, 2552333, 3493499, 3447281, 7455438, 1748304, 7613479, 1582964, 9353610, 4161262, 9341707, 1374920, 3612393, 5155326, 7612309, 1751012, 3344169, 8669800, 9659720, 1952127, 9671349, 2177852, 1215837, 9190527, 1710230, 6140015, 5378238, 1304308, 7710459, 6822402, 1846846, 5634309, 9728197, 9857045, 5266151, 1747512, 6434565, 4980944, 7975735, 2968600, 2414939, 1337052, 3437904, 7281635, 7365460, 2501597, 9653175, 7231323, 8592064, 8510882, 9600628, 8269404, 8536971, 1388562, 3911760, 2890581, 4208818, 6165146, 9994782, 4382118, 2345465, 5233495, 8917172, 4568384, 9961273, 2864425, 6480452, 2545940, 5202716, 5932389, 4852609, 5664695, 7971805}
	arr2 := []int{8299740, 9361730, 4412913, 4784698, 7307894, 4887723, 4656373, 6851818, 4585883, 6267112, 1195791, 2432490, 2218541, 8592481, 2327155, 3853490, 1617246, 8615151, 1344525, 4802420, 6580598, 6378162, 4672373, 2886671, 3958496, 6324197, 1601967, 7421870, 7354689, 6845557, 3634049, 4812776, 8843319, 3104985, 2188282, 8284876, 5341579, 7901226, 4583778, 3848127, 4641464, 9861055, 2648267, 8661757, 9770194, 6641170, 3619502, 7578451, 3143098, 3102908, 7245034, 7637564, 4351780, 9285795, 5499287, 4906118, 1366815, 5859901, 6765375, 3381606, 9682030, 6254567, 3417403, 2593282, 6908692, 9779429, 7844164, 6454609, 6648065, 9827826, 6230304, 4363589, 9661167, 2715387, 4665052, 8946340, 2821829, 7117913, 8771306, 8746510, 7274480, 8179938, 4197772, 3787690, 4779209, 8986845, 3842349, 6230460, 1210986, 2928858, 5322721, 2469091, 5457910, 7370179, 3965982, 5536633, 3589718, 7718320, 5488602, 5921443, 7715022, 3710652, 1789532, 1707985, 4711213, 4137639, 9620964, 6530545, 1434972, 7274437, 5316799, 8397478, 1305898, 9321932, 9411472, 3512328, 2995835, 1771567, 6174526, 6845480, 5912358, 5695872, 4572976, 6284474, 5163324}

	result := longestCommonPrefix(arr1, arr2)
	fmt.Println("The length of the longest common prefix is:", result)
}
