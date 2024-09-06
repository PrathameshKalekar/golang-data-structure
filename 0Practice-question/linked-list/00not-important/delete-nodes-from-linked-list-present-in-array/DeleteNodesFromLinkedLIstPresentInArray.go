/*
3217 ou are given an array of integers nums and the head of a linked list. Return the head of the modified linked list after removing all nodes from the linked list that have a value that exists in nums.

Example 1:

Input: nums = [1,2,3], head = [1,2,3,4,5]

Output: [4,5]

Explanation:

Remove the nodes with values 1, 2, and 3.

Example 2:

Input: nums = [1], head = [1,2,1,2,1,2]

Output: [2,2,2]

Explanation:

Remove the nodes with value 1.

Example 3:

Input: nums = [5], head = [1,2,3,4]

Output: [1,2,3,4]

Explanation:

No node has value 5.
*/
package main

import "fmt"

func main() {
	node := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}}
	node.Display()
	answer := deleteNodesFromLinkedListPresentInArray([]int{3, 2}, node)

	answer.Display()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) Display() {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
	fmt.Println()
}

func deleteNodesFromLinkedListPresentInArray(arr []int, head *ListNode) *ListNode {
	nodeMap := make(map[int]bool)
	for _, value := range arr {
		nodeMap[value] = true
	}
	result := &ListNode{}
	dummy := result
	for head != nil {
		if !nodeMap[head.Val] {
			dummy.Next = head

			dummy = dummy.Next
		}
		if head.Next == nil {
			dummy.Next = nil
		}
		head = head.Next
	}
	return result.Next
}
