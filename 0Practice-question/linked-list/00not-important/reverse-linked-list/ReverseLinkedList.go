/*
206 Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:

Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
Example 2:

Input: head = [1,2]
Output: [2,1]
Example 3:

Input: head = []
Output: []
*/
package main

func main() {

}

type ListNode struct {
	Val     int
	NexNode *ListNode
}

func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil || head.NexNode == nil {
		return head
	}
	//temp array to keep track reversse order of the linked list
	temp := []int{}
	for head != nil {
		//append each value in array
		temp = append(temp, head.Val)
		head = head.NexNode
	}
	//create a result to add the reverse order on it
	result := &ListNode{}
	dummy := result
	//from the reverse order of the array add the reverse order linked list
	for i := len(temp) - 1; i >= 0; i-- {
		dummy.NexNode = &ListNode{Val: temp[i]}
		dummy = dummy.NexNode
	}
	//return the linked list .next cause current listNode is empty
	return result.NexNode
}
