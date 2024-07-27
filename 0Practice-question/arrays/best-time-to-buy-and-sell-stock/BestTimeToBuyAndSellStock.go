/*
121 You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
*/
package main

import "fmt"

func main() {
	stockPrices := []int{3, 5, 2, 6, 7, 5, 3, 1}
	answer := bestTimeToBuyAndSellStock(stockPrices)
	fmt.Printf("Max profit after buy and sell is : %d", answer)
}

func bestTimeToBuyAndSellStock(stockPrices []int) int {
	low := stockPrices[0]
	maxProfit := 0
	//iterating through each value in the stick price
	for i := 0; i < len(stockPrices); i++ {
		//if current value is less that low means its a best time to buy so low = i th term
		if stockPrices[i] < low {
			low = stockPrices[i]
		} else {
			//else after buy time to sell
			price := stockPrices[i] - low
			//if the selling price is greter than max profit best time to sell
			if price > maxProfit {
				maxProfit = price
			}
		}
	}
	//return max profit
	return maxProfit
}
