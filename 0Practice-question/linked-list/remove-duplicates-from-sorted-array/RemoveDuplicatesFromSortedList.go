/*
83 Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

Example 1:

Input: head = [1,1,2]
Output: [1,2]
Example 2:

Input: head = [1,1,2,3,3]
Output: [1,2,3]
*/
package main

import "fmt"

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 6}}}}}}}
	node.Display()
	answer := removeDuplicatesFromSortedList(node)
	answer.Display()
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicatesFromSortedList(head *ListNode) *ListNode {
	current := head

	// Traverse through the list
	for current != nil && current.Next != nil {
		// If current node is equal to the next node, skip the next node
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			// Otherwise, move to the next node
			current = current.Next
		}
	}

	return head
}

func (list *ListNode) Display() {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
	fmt.Println()
}
