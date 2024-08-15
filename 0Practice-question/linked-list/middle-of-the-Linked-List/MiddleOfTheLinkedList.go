/*
876 Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.

Example 1:

Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.
Example 2:

Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
*/
package main

import "fmt"

func main() {
	// Creating a linked list: 1 -> 2 -> 3 -> 4 -> 5 -> nil
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 1, Next: node2}

	middle := middleOfLinkedList(head)
	if middle != nil {
		fmt.Printf("The middle node value is: %d\n", middle.Val)
	} else {
		fmt.Println("The list is empty.")
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleOfLinkedList(head *ListNode) *ListNode {
	slow, fast := head, head
	//find the middle in linked list
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//fast goes two step therefore slow is on middle when fast reach the end
	return slow
}
