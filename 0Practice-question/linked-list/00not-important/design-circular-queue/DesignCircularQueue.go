/*
641 Design your implementation of the circular double-ended queue (deque).

Implement the MyCircularDeque class:

MyCircularDeque(int k) Initializes the deque with a maximum size of k.
boolean insertFront() Adds an item at the front of Deque. Returns true if the operation is successful, or false otherwise.
boolean insertLast() Adds an item at the rear of Deque. Returns true if the operation is successful, or false otherwise.
boolean deleteFront() Deletes an item from the front of Deque. Returns true if the operation is successful, or false otherwise.
boolean deleteLast() Deletes an item from the rear of Deque. Returns true if the operation is successful, or false otherwise.
int getFront() Returns the front item from the Deque. Returns -1 if the deque is empty.
int getRear() Returns the last item from Deque. Returns -1 if the deque is empty.
boolean isEmpty() Returns true if the deque is empty, or false otherwise.
boolean isFull() Returns true if the deque is full, or false otherwise.

Example 1:

Input
["MyCircularDeque", "insertLast", "insertLast", "insertFront", "insertFront", "getRear", "isFull", "deleteLast", "insertFront", "getFront"]
[[3], [1], [2], [3], [4], [], [], [], [4], []]
Output
[null, true, true, true, false, 2, true, true, true, 4]

Explanation
MyCircularDeque myCircularDeque = new MyCircularDeque(3);
myCircularDeque.insertLast(1);  // return True
myCircularDeque.insertLast(2);  // return True
myCircularDeque.insertFront(3); // return True
myCircularDeque.insertFront(4); // return False, the queue is full.
myCircularDeque.getRear();      // return 2
myCircularDeque.isFull();       // return True
myCircularDeque.deleteLast();   // return True
myCircularDeque.insertFront(4); // return True
myCircularDeque.getFront();     // return 4
*/
package main

import "fmt"

type MyCircularDeque struct {
	data     []int
	size     int
	capacity int
	front    int
	rear     int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		data:     make([]int, k+1), // Circular deque uses size k+1 to differentiate between empty and full states.
		size:     0,
		capacity: k + 1, // We use one extra space to differentiate full and empty.
		front:    0,
		rear:     0,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.front = (this.front - 1 + this.capacity) % this.capacity // Move front pointer backward
	this.data[this.front] = value
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.rear] = value
	this.rear = (this.rear + 1) % this.capacity // Move rear pointer forward
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % this.capacity // Move front pointer forward
	this.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.rear = (this.rear - 1 + this.capacity) % this.capacity // Move rear pointer backward
	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.rear-1+this.capacity)%this.capacity]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.capacity-1
}

func main() {
	circularDeque := Constructor(3)
	fmt.Println(circularDeque.InsertLast(1))  // returns true
	fmt.Println(circularDeque.InsertLast(2))  // returns true
	fmt.Println(circularDeque.InsertFront(3)) // returns true
	fmt.Println(circularDeque.InsertFront(4)) // returns false, queue is full
	fmt.Println(circularDeque.GetRear())      // returns 2
	fmt.Println(circularDeque.IsFull())       // returns true
	fmt.Println(circularDeque.DeleteLast())   // returns true
	fmt.Println(circularDeque.InsertFront(4)) // returns true
	fmt.Println(circularDeque.GetFront())     // returns 4
}
