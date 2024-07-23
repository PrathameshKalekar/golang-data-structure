package main

import "fmt"

//creating a queue structure
type Queue struct {
	content []int
}

func main() {
	//creating queue instance
	var queue Queue

	queue.Add(3)
	queue.Add(6)
	queue.Add(9)
	queue.Add(2)
	queue.Display()
	queue.Pop()
	queue.Display()
	fmt.Println(queue.FirstAddedItem())
	fmt.Println(queue.LastAddedItem())
}

//struct method to add elements in queue
func (queue *Queue) Add(item int) {
	queue.content = append(queue.content, item)
}

//struct method to read elelmets ini queue
func (queue *Queue) Display() {
	fmt.Println(queue.content)
}

//struct method to pop the first element added
func (queue *Queue) Pop() (int, error) {
	if len(queue.content) == 0 {
		return 0, fmt.Errorf("No content found")
	}
	firstItem := queue.content[0]
	queue.content = queue.content[1:]
	return firstItem, nil
}

//struct method to display first item added in queue
func (queue *Queue) FirstAddedItem() (int, error) {
	if len(queue.content) == 0 {
		return 0, fmt.Errorf("No content found")
	}
	return queue.content[0], nil
}

//struct method to display last item added in queue
func (queue *Queue) LastAddedItem() (int, error) {
	if len(queue.content) == 0 {
		return 0, fmt.Errorf("No content found")
	}
	return queue.content[len(queue.content)-1], nil
}
