package main

import "fmt"

//To create a node for each linked list element
type Node struct {
	Value    int
	NextNode *Node
}

//Circular linked list structure created having a head node
type CircularLinkedList struct {
	Head *Node
}

func main() {
	//creating a instance of linked list
	var list CircularLinkedList
	list.Add(3)
	list.Add(5)
	list.Add(2)
	list.Add(1)
	list.Display()
}

//struct method to add items in circular linked list
func (list *CircularLinkedList) Add(value int) {
	newNode := &Node{
		Value: value,
	}

	if list.Head == nil {
		list.Head = newNode
		newNode.NextNode = list.Head
	} else {
		current := list.Head
		for current.NextNode != list.Head {
			current = current.NextNode
		}
		current.NextNode = newNode
		newNode.NextNode = list.Head
	}
}

//struct method to display items in circular linked list
func (list *CircularLinkedList) Display() {
	current := list.Head
	if current == nil {
		fmt.Println("NO Node Found")
	} else {
		fmt.Print("Circular Linked List : ")
		for {
			fmt.Printf("%d ", current.Value)
			current = current.NextNode
			if current.NextNode == list.Head {
				fmt.Printf("%d ", current.Value)
				break
			}
		}
	}
}
