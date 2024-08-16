/*
86 Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example 1:

Input: head = [1,4,3,2,5,2], x = 3
Output: [1,2,2,4,3,5]
Example 2:

Input: head = [2,1], x = 2
Output: [1,2]
*/
package main

import "fmt"

func main() {
	node := &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}
	answer := partitionList(node, 3)
	answer.Display()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partitionList(head *ListNode, x int) *ListNode {
	// Two dummy nodes to partition the list into two parts
	left, right := &ListNode{}, &ListNode{}
	dummyLeft, dummyRight := left, right

	// Traverse the list and partition nodes
	for head != nil {
		if head.Val < x {
			dummyLeft.Next = head
			dummyLeft = dummyLeft.Next
		} else {
			dummyRight.Next = head
			dummyRight = dummyRight.Next
		}
		head = head.Next
	}

	// End the right list
	dummyRight.Next = nil
	// Connect the left list to the right list
	dummyLeft.Next = right.Next

	return left.Next
}

func (list *ListNode) Display() {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
	fmt.Println()
}
