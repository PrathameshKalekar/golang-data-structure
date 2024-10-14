/*
2530 You are given a 0-indexed integer array nums and an integer k. You have a starting score of 0.

In one operation:

choose an index i such that 0 <= i < nums.length,
increase your score by nums[i], and
replace nums[i] with ceil(nums[i] / 3).
Return the maximum possible score you can attain after applying exactly k operations.

The ceiling function ceil(val) is the least integer greater than or equal to val.

Example 1:

Input: nums = [10,10,10,10,10], k = 5
Output: 50
Explanation: Apply the operation to each array element exactly once. The final score is 10 + 10 + 10 + 10 + 10 = 50.
Example 2:

Input: nums = [1,10,3,3,3], k = 3
Output: 17
Explanation: You can do the following operations:
Operation 1: Select i = 1, so nums becomes [1,4,3,3,3]. Your score increases by 10.
Operation 2: Select i = 1, so nums becomes [1,2,3,3,3]. Your score increases by 4.
Operation 3: Select i = 2, so nums becomes [1,1,1,3,3]. Your score increases by 3.
The final score is 10 + 4 + 3 = 17.
*/
package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	arr := []int{10, 10, 10, 10, 10}
	answer := maximalScoreAfterApplyingKOperations(arr, 5)
	fmt.Println(answer)
}

// Priority queue (max-heap) implementation
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-Heap: largest element first
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Function to calculate the maximum score after k operations
func maximalScoreAfterApplyingKOperations(nums []int, k int) int64 {
	// Initialize the max heap
	h := &MaxHeap{}
	heap.Init(h)

	// Push all elements into the heap
	for _, num := range nums {
		heap.Push(h, num)
	}

	var ans int64 = 0

	// Perform k operations
	for i := 0; i < k; i++ {
		// Pop the largest element
		val := heap.Pop(h).(int)
		// Add the largest element to the answer
		ans += int64(val)
		// Push the reduced value back into the heap (ceil(val / 3))
		heap.Push(h, int(math.Ceil(float64(val)/3.0)))
	}

	return ans
}
