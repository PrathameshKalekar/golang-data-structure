/*
21 You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:

Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]
*/
package main

import "fmt"

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	answer := mergeTwoSortedLists(list1, list2)
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

func mergeTwoSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	newList := result
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			newList.Next = &ListNode{Val: list1.Val}
			newList = newList.Next
			list1 = list1.Next
		} else {
			newList.Next = &ListNode{Val: list2.Val}
			newList = newList.Next
			list2 = list2.Next
		}
	}
	for list1 != nil {
		newList.Next = &ListNode{Val: list1.Val}
		newList = newList.Next
		list1 = list1.Next
	}
	for list2 != nil {
		newList.Next = &ListNode{Val: list2.Val}
		newList = newList.Next
		list2 = list2.Next
	}

	return result.Next
}
