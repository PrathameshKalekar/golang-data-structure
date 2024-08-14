/*
24 Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

Example 1:

Input: head = [1,2,3,4]
Output: [2,1,4,3]
Example 2:

Input: head = []
Output: []
Example 3:

Input: head = [1]
Output: [1]
*/
package main

import "fmt"

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7}}}}}
	fmt.Print("Before swap : ")
	node.Display()
	answer := swapNodesInPairs(node)
	fmt.Print("after swap : ")
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
func swapNodesInPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	current := dummy

	// Loop through the list in pairs
	for current.Next != nil && current.Next.Next != nil {
		// Initialize the nodes to swap
		first := current.Next
		second := current.Next.Next

		// Swap the nodes
		first.Next = second.Next
		second.Next = first
		current.Next = second

		// Move to the next pair
		current = first
	}

	return dummy.Next
}

// func swapNodesInPairs(head *ListNode) *ListNode {
// 	//if head is empty or head have only onoe node then return the node
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	result := &ListNode{}
// 	dummy := result
// 	temp := []int{}
// 	for head != nil {
// 		temp = append(temp, head.Val)
// 		if len(temp) == 2 {
// 			for i := 1; i >= 0; i-- {
// 				dummy.Next = &ListNode{Val: temp[i]}
// 				dummy = dummy.Next
// 			}
// 			temp = []int{}
// 		}
// 		head = head.Next
// 	}
// 	if len(temp) > 0 {
// 		dummy.Next = &ListNode{Val: temp[0]}
// 	}
// 	return result.Next
// }
