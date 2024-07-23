package main

import "fmt"

//To create a node for each linked list element
type Node struct {
	Value    int
	NextNode *Node
}

//Linked list structure created having a head node
type LinkedList struct {
	Head *Node
}

func main() {
	//creating a instance of linked list
	var list LinkedList

	list.Add(3)
	list.Add(34)
	list.Add(543)
	list.Display()
}

//struct method to add items in linked list
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

//struct method to display linked list
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
