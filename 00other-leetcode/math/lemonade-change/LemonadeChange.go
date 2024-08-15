/*
860 At a lemonade stand, each lemonade costs $5. Customers are standing in a queue to buy from you and order one at a time (in the order specified by bills). Each customer will only buy one lemonade and pay with either a $5, $10, or $20 bill. You must provide the correct change to each customer so that the net transaction is that the customer pays $5.

Note that you do not have any change in hand at first.

Given an integer array bills where bills[i] is the bill the ith customer pays, return true if you can provide every customer with the correct change, or false otherwise.

Example 1:

Input: bills = [5,5,5,10,20]
Output: true
Explanation:
From the first 3 customers, we collect three $5 bills in order.
From the fourth customer, we collect a $10 bill and give back a $5.
From the fifth customer, we give a $10 bill and a $5 bill.
Since all customers got correct change, we output true.
Example 2:

Input: bills = [5,5,10,10,20]
Output: false
Explanation:
From the first two customers in order, we collect two $5 bills.
For the next two customers in order, we collect a $10 bill and give back a $5 bill.
For the last customer, we can not give the change of $15 back because we only have two $10 bills.
Since not every customer received the correct change, the answer is false.
*/
package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 5, 10, 20}
	answer := lemonadeChange(arr)
	fmt.Println(answer)
}

func lemonadeChange(arr []int) bool {
	//for managing the five and ten coins
	five, ten := 0, 0
	//for each coustomer
	for _, bill := range arr {
		//if coustemer gives 5 $ then add five
		if bill == 5 {
			five++
		} else if bill == 10 {
			//if customer give 10$ then give 5 doller change if we dont have 5$ then return flase
			if five > 0 {
				five--
				ten++
			} else {
				return false
			}
		} else if bill == 20 {
			//if customer give 20$ then give 5 ad 10 doller change or 3 change of 5 doller if we dont have any then return flase
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five > 2 {
				five -= 3
			} else {
				return false
			}
		}
	}
	//if all customers clear means all changes are sucess return true
	return true
}
