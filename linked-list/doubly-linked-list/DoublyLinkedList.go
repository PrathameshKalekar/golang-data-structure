package main

import "fmt"

//To create a node for each linked list element
type Node struct {
	Value    int
	PrevNode *Node
	NextNode *Node
}

//Doubly linked list structure created having a head node
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func main() {
	//creating a instance of linked list
	var list DoublyLinkedList

	list.Add(4)
	list.Add(3)
	list.Add(2)
	list.Add(1)
	list.ReadForward()
	list.ReadBackward()
}

//struct method to add items in doubly linked list
func (list *DoublyLinkedList) Add(value int) {
	newNode := &Node{
		Value: value,
	}

	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.PrevNode = list.Tail
		list.Tail.NextNode = newNode
		list.Tail = newNode
	}
}

//struct method to read element in forward form
func (list *DoublyLinkedList) ReadForward() {
	current := list.Head
	fmt.Print("Reading forward : ")
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.NextNode
	}
}

//struct method to read element in backward form
func (list *DoublyLinkedList) ReadBackward() {
	current := list.Tail
	fmt.Print("Reading backward : ")
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.PrevNode
	}
}
