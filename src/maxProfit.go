package main

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3287/
import "fmt"

func main() {
	fmt.Println("Hello")

	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(prices)

	fmt.Println(profit)
}

func maxProfit(prices []int) int {
	var maxprofit int
	length := len(prices)
	for i := 1; i < length; i++ {
		if prices[i] > prices[i-1] {
			maxprofit += prices[i] - prices[i-1]
		}
	}
	return maxprofit
}
