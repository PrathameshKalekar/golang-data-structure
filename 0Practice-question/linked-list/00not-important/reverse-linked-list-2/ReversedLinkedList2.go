/*
92 Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.

Example 1:

Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]
Example 2:

Input: head = [5], left = 1, right = 1
Output: [5]
*/
package main

import "fmt"

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}
	answer := reverseLinkedList2(node, 2, 4)
	answer.Display()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLinkedList2(head *ListNode, left, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	// Create a dummy node to simplify edge cases
	dummy := &ListNode{Next: head}
	prev := dummy

	// Move `prev` to one node before the position `m`
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// Start of the sublist to reverse
	start := prev.Next
	// The node that will be after the reversed sublist
	then := start.Next

	// Reverse the sublist between m and n
	for i := 0; i < right-left; i++ {
		start.Next = then.Next
		then.Next = prev.Next
		prev.Next = then
		then = start.Next
	}

	return dummy.Next
}

//for displaying the node
func (list *ListNode) Display() {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
	fmt.Println()
}
