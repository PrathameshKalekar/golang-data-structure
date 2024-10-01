/*
1497 Given an array of integers arr of even length n and an integer k.

We want to divide the array into exactly n / 2 pairs such that the sum of each pair is divisible by k.

Return true If you can find a way to do that or false otherwise.

Example 1:

Input: arr = [1,2,3,4,5,10,6,7,8,9], k = 5
Output: true
Explanation: Pairs are (1,9),(2,8),(3,7),(4,6) and (5,10).
Example 2:

Input: arr = [1,2,3,4,5,6], k = 7
Output: true
Explanation: Pairs are (1,6),(2,5) and(3,4).
Example 3:

Input: arr = [1,2,3,4,5,6], k = 10
Output: false
Explanation: You can try all possible pairs to see that there is no way to divide arr into 3 pairs each with sum divisible by 10.
*/
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}
	fmt.Println(checkIfArrayPairsAreDivisibleByK(arr, 5))
}

func checkIfArrayPairsAreDivisibleByK(arr []int, k int) bool {
	mods := make(map[int]int)
	for _, v := range arr {
		mods[((v%k)+k)%k]++
	}
	if mods[0]%2 == 1 {
		return false
	}
	for i := 1; i < k; i++ {
		if mods[i] != mods[k-i] {
			return false
		}
	}
	return true
}
