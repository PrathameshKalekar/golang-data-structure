package main

import "fmt"

type Node struct {
	Value    int
	NextNode *Node
}

type LinkedList struct {
	Head *Node
}

func main() {
	var list LinkedList
	list.Add(3)
	list.Add(34)
	list.Add(543)
	list.Display()
}

func (list *LinkedList) Add(value int) {
	newNode := &Node{
		Value: value,
	}

	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		for current.NextNode != nil {
			current = current.NextNode
		}
		current.NextNode = newNode
	}
}

func (list *LinkedList) Display() {
	current := list.Head
	if current == nil {
		fmt.Println("No Node Found")
	} else {
		for current != nil {
			fmt.Printf("%d ", current.Value)
			current = current.NextNode
		}
	}
}
