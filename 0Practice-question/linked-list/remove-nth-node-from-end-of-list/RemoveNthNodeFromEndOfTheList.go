/*
19 Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:

Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Example 2:

Input: head = [1], n = 1
Output: []
Example 3:

Input: head = [1,2], n = 1
Output: [1]
*/
package main

import (
	"fmt"
)

func main() {
	list := &ListNode{Val: 1, NextNode: &ListNode{Val: 2}}
	result := removeNthNodeFromEndOfTheList(list, 1)
	result.Display()
}

type ListNode struct {
	Val      int
	NextNode *ListNode
}

func removeNthNodeFromEndOfTheList(list *ListNode, index int) *ListNode {
	result := &ListNode{}
	dummy := result
	arr := []int{}
	//we cant go backward in list so add a list into array
	for list != nil {
		arr = append(arr, list.Val)
		list = list.NextNode
	}
	//for length of array
	for i := 0; i < len(arr); i++ {
		// if the i == len(arr)-index  means the current index is n from the last so we dont want it in listNode sp skip it
		if i == len(arr)-index {
			continue
		}
		//its a pointer so we cannot play with direcly result so we created a dummy that we can play with
		dummy.NextNode = &ListNode{Val: arr[i]}
		dummy = dummy.NextNode
	}
	//current node in list is empty so return result.NextNOde
	return result.NextNode
}

func (list *ListNode) Display() {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.NextNode
	}
	fmt.Println()
}
