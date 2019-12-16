package main

/* implementation of the Share Prices problem */

import (
	"fmt"
)

func getMaxProfit(prices []int, n int) []int {
	min := make([][]int, n)
	days := make([]int, 2)
	minPrice := prices[0]
	minPriceDay := 0

	for i := range min {
		min[i] = make([]int, 2)
	}

	min[0][0] = prices[0]
	for i := 1; i < n; i++ {
		min[i][0] = minPrice
		min[i][1] = minPriceDay
		if minPrice > prices[i] {
			minPrice = prices[i]
			minPriceDay = i
		}
	}

	profit := 0
	for i := 1; i < n; i++ {
		if profit < prices[i]-min[i][0] {
			profit = prices[i] - min[i][0]
			days[0] = min[i][1]
			days[1] = i
		}
	}

	return days
}

func main() {
	fmt.Println("Please enter the number of days, ")
	fmt.Println("please then enter the prices of each day :-")

	var n int
	fmt.Scanln(&n)
	prices := make([]int, n)
	for i := range prices {
		fmt.Scan(&prices[i])
	}

	days := getMaxProfit(prices, n)
	fmt.Println("Buy on :", days[0])
	fmt.Println("Sell on :", days[1])
}
