/*
203 Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

Example 1:

Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]
Example 2:

Input: head = [], val = 1
Output: []
Example 3:

Input: head = [7,7,7,7], val = 7
Output: []
*/
package main

import "fmt"

func main() {
	node := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}}
	answer := removeLinkedListElement(node, 3)
	answer.Display()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeLinkedListElement(head *ListNode, val int) *ListNode {
	// Handle cases where the head itself needs to be removed
	for head != nil && head.Val == val {
		head = head.Next
	}

	// If the entire list becomes empty, return nil
	if head == nil {
		return nil
	}

	dummy := head
	for dummy.Next != nil {
		if dummy.Next.Val == val {
			// Remove the node by skipping it
			dummy.Next = dummy.Next.Next
		} else {
			dummy = dummy.Next
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
