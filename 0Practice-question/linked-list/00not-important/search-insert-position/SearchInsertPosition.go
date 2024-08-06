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
	// Creating a list with no cycle
	node1 := &ListNode{Val: 14}
	node2 := &ListNode{Val: 14}
	node3 := &ListNode{Val: 14}
	node4 := &ListNode{Val: 14}
	node5 := &ListNode{Val: 14}
	node6 := &ListNode{Val: 14}
	node7 := &ListNode{Val: 14}

	node1.NextNode = node2
	node2.NextNode = node3
	node3.NextNode = node4
	node4.NextNode = node5
	node5.NextNode = node6
	node6.NextNode = node7
	node7.NextNode = nil // No cycle here

	fmt.Println(searchInsertPosition(node1)) // Output: false

	// Creating a list with a cycle for testing
	node1 = &ListNode{Val: 14}
	node2 = &ListNode{Val: 14}
	node3 = &ListNode{Val: 14}
	node4 = &ListNode{Val: 14}
	node5 = &ListNode{Val: 14}
	node6 = &ListNode{Val: 14}
	node7 = &ListNode{Val: 14}

	node1.NextNode = node2
	node2.NextNode = node3
	node3.NextNode = node4
	node4.NextNode = node5
	node5.NextNode = node6
	node6.NextNode = node7
	node7.NextNode = node3 // Cycle here

	fmt.Println(searchInsertPosition(node1)) // Output: true
}

type ListNode struct {
	Val      int
	NextNode *ListNode
}

func searchInsertPosition(head *ListNode) bool {
	if head == nil || head.NextNode == nil {
		return false
	}
	start, end := head, head
	for end != nil && end.NextNode != nil {
		start = start.NextNode
		end = end.NextNode.NextNode
		if start == end {
			return true
		}
	}
	return false
}
