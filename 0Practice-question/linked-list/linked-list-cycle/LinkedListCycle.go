/*
141 Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

Example 1:

Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
Example 2:

Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
Example 3:

Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
*/
package main

import "fmt"

func main() {
	// Manually creating a linked list: 1 -> 2 -> 3 -> 4 -> 5 -> nil
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 1, Next: node2}

	// Case 1: No cycle in the list
	fmt.Println("Linked List 1 (without cycle):")

	if linkedListCycle(head) {
		fmt.Println("Cycle detected.")
	} else {
		fmt.Println("No cycle detected.")
	}

	// Case 2: Manually creating a cycle in the list at position 1 (node2)
	fmt.Println("\nManually creating a cycle...")
	node5.Next = node2 // Manually linking node 5 back to node 2 to create a cycle

	// Testing after creating the cycle
	fmt.Println("Linked List 2 (with a cycle):")
	if linkedListCycle(head) {
		fmt.Println("Cycle detected.")
	} else {
		fmt.Println("No cycle detected.")
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func linkedListCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	start, end := head, head
	for end != nil && end.Next != nil {
		start = start.Next
		end = end.Next.Next
		if start == end {
			return true
		}
	}
	return false
}
