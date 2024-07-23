package main

import "fmt"

//creating a stack structure
type Stack struct {
	content []int
}

func main() {
	//creating a stack instance
	var stack Stack

	stack.Add(2)
	stack.Add(4)
	stack.Display()
	stack.Pop()
	stack.Display()
	fmt.Println(stack.FirstAddedItem())
	fmt.Println(stack.LastAddedItem())
}

//struct method to add item in stack
func (stack *Stack) Add(item int) {
	stack.content = append(stack.content, item)
}

//struct method to display items in stack
func (stack *Stack) Display() {
	fmt.Println(stack.content)
}

//struct method to pop last added item in the stack
func (stack *Stack) Pop() (int, error) {
	if len(stack.content) == 0 {
		return 0, fmt.Errorf("No content Found")
	}
	lastItemIndex := len(stack.content) - 1
	lastItem := stack.content[lastItemIndex]

	stack.content = stack.content[:lastItemIndex]
	return lastItem, nil
}

//struct method to check first added item
func (stack *Stack) FirstAddedItem() (int, error) {
	if len(stack.content) == 0 {
		return 0, fmt.Errorf("No Content Found")
	}

	return stack.content[0], nil
}

//struct method to check last added item
func (stack *Stack) LastAddedItem() (int, error) {
	if len(stack.content) == 0 {
		return 0, fmt.Errorf("No Content Found")
	}

	lastItem := stack.content[len(stack.content)-1]
	return lastItem, nil
}
