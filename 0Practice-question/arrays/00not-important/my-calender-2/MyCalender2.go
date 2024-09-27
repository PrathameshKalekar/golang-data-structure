/*
731 You are implementing a program to use as your calendar. We can add a new event if adding the event will not cause a triple booking.

A triple booking happens when three events have some non-empty intersection (i.e., some moment is common to all the three events.).

The event can be represented as a pair of integers start and end that represents a booking on the half-open interval [start, end), the range of real numbers x such that start <= x < end.

Implement the MyCalendarTwo class:

MyCalendarTwo() Initializes the calendar object.
boolean book(int start, int end) Returns true if the event can be added to the calendar successfully without causing a triple booking. Otherwise, return false and do not add the event to the calendar.

Example 1:

Input
["MyCalendarTwo", "book", "book", "book", "book", "book", "book"]
[[], [10, 20], [50, 60], [10, 40], [5, 15], [5, 10], [25, 55]]
Output
[null, true, true, true, false, true, true]

Explanation
MyCalendarTwo myCalendarTwo = new MyCalendarTwo();
myCalendarTwo.book(10, 20); // return True, The event can be booked.
myCalendarTwo.book(50, 60); // return True, The event can be booked.
myCalendarTwo.book(10, 40); // return True, The event can be double booked.
myCalendarTwo.book(5, 15);  // return False, The event cannot be booked, because it would result in a triple booking.
myCalendarTwo.book(5, 10); // return True, The event can be booked, as it does not use time 10 which is already double booked.
myCalendarTwo.book(25, 55); // return True, The event can be booked, as the time in [25, 40) will be double booked with the third event, the time [40, 50) will be single booked, and the time [50, 55) will be double booked with the second event.
*/
package main

import "fmt"

// MyCalendarTwo stores the single and double bookings.
type MyCalendarTwo struct {
	singles [][]int
	doubles [][]int
}

// Constructor initializes the MyCalendarTwo object.
func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		singles: [][]int{},
		doubles: [][]int{},
	}
}

// book checks if the event can be booked without causing a triple booking.
func (this *MyCalendarTwo) Book(start int, end int) bool {
	// Check if the event overlaps with any double bookings.
	for _, double := range this.doubles {
		if max(start, double[0]) < min(end, double[1]) {
			// Triple booking would happen, so return false.
			return false
		}
	}

	// Check for overlaps with single bookings and update double bookings.
	for _, single := range this.singles {
		if max(start, single[0]) < min(end, single[1]) {
			// There is an overlap, so add it to the double bookings.
			this.doubles = append(this.doubles, []int{max(start, single[0]), min(end, single[1])})
		}
	}

	// Add the new event to the list of single bookings.
	this.singles = append(this.singles, []int{start, end})
	return true
}

// Helper functions to calculate min and max.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Main function to test the implementation.
func main() {
	myCalendarTwo := Constructor()
	fmt.Println(myCalendarTwo.Book(10, 20)) // Output: true
	fmt.Println(myCalendarTwo.Book(50, 60)) // Output: true
	fmt.Println(myCalendarTwo.Book(10, 40)) // Output: true
	fmt.Println(myCalendarTwo.Book(5, 15))  // Output: false
	fmt.Println(myCalendarTwo.Book(5, 10))  // Output: true
	fmt.Println(myCalendarTwo.Book(25, 55)) // Output: true
}
