/*
82 Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.

Example 1:

Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]
Example 2:

Input: head = [1,1,1,2,3]
Output: [2,3]
*/
package main

import "fmt"

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}
	answer := removeDuplicatesFromSortedList(node)
	answer.Display()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicatesFromSortedList(head *ListNode) *ListNode {
	// Dummy node to handle edge cases easily
	dummy := &ListNode{Next: head}
	prev := dummy // Prev will track the node before the current group of duplicates

	for head != nil {
		// Check if the current node has duplicates
		if head.Next != nil && head.Val == head.Next.Val {
			// Skip all nodes with the same value
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			// Link prev to the node after the last duplicate
			prev.Next = head.Next
		} else {
			// No duplicates, move prev forward
			prev = prev.Next
		}
		// Move head forward
		head = head.Next
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
