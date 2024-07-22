package main

import "fmt"

type Node struct {
	Value    int
	PrevNode *Node
	NextNode *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func main() {
	var list DoublyLinkedList

	list.Add(4)
	list.Add(3)
	list.Add(2)
	list.Add(1)
	list.ReadForward()
	list.ReadBackward()
}

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

func (list *DoublyLinkedList) ReadForward() {
	current := list.Head
	fmt.Print("Reading forward : ")
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.NextNode
	}
}

func (list *DoublyLinkedList) ReadBackward() {
	current := list.Tail
	fmt.Print("Reading backward : ")
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.PrevNode
	}
}
