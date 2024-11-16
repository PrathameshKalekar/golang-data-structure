package main

import "fmt"

// Function to find the power of k-size subarrays
func findThePowerOfKSizeSubarray(nums []int, k int) []int {
	n := len(nums)
	result := make([]int, 0, n-k+1)

	// Traverse through the array and process each subarray of size k
	for i := 0; i <= n-k; i++ {
		isConsecutive := true
		// Check if the subarray is sorted and consecutive
		for j := 1; j < k; j++ {
			if nums[i+j] != nums[i+j-1]+1 {
				isConsecutive = false
				break
			}
		}

		if isConsecutive {
			// Append the maximum element if sorted and consecutive
			result = append(result, nums[i+k-1])
		} else {
			// Append -1 otherwise
			result = append(result, -1)
		}
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 3, 4, 3, 2, 5}
	k1 := 3
	fmt.Println(findThePowerOfKSizeSubarray(nums1, k1)) // Output: [3, 4, -1, -1, -1]

	nums2 := []int{2, 2, 2, 2, 2}
	k2 := 4
	fmt.Println(findThePowerOfKSizeSubarray(nums2, k2)) // Output: [-1, -1]

	nums3 := []int{3, 2, 3, 2, 3, 2}
	k3 := 2
	fmt.Println(findThePowerOfKSizeSubarray(nums3, k3)) // Output: [-1, 3, -1, 3, -1]
}
