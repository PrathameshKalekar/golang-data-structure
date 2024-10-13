/*
632 You have k lists of sorted integers in non-decreasing order. Find the smallest range that includes at least one number from each of the k lists.

We define the range [a, b] is smaller than range [c, d] if b - a < d - c or a < c if b - a == d - c.

Example 1:

Input: nums = [[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]
Output: [20,24]
Explanation:
List 1: [4, 10, 15, 24,26], 24 is in range [20,24].
List 2: [0, 9, 12, 20], 20 is in range [20,24].
List 3: [5, 18, 22, 30], 22 is in range [20,24].
Example 2:

Input: nums = [[1,2,3],[1,2,3],[1,2,3]]
Output: [1,1]
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}
	answer := smallestRangeCoveringElementsFromKSortedLists(nums)
	fmt.Println(answer)
}

type data struct {
	num   int
	group int
}

func smallestRangeCoveringElementsFromKSortedLists(nums [][]int) []int {
	var list []data
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			list = append(list, data{nums[i][j], i})
		}
	}

	sort.Slice(list, func(i int, j int) bool {
		if list[i].num == list[j].num {
			return list[i].group < list[j].group
		}
		return list[i].num < list[j].num
	})

	i, j := 0, 0
	dict := make(map[int]int)
	cnt := 0
	res1, res2 := 0, math.MaxInt64
	for j < len(list) {
		if dict[list[j].group] == 0 {
			cnt++
		}
		dict[list[j].group]++
		j++

		for cnt == len(nums) {
			if res2-res1 > list[j-1].num-list[i].num {
				res1, res2 = list[i].num, list[j-1].num
			}
			if dict[list[i].group] == 1 {
				cnt--
			}
			dict[list[i].group]--
			i++
		}
	}
	return []int{res1, res2}
}
