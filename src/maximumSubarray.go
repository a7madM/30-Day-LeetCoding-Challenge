package main

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3285/

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	result := maxSubArray(nums)
	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	} else if length == 1 {
		return nums[0]
	}

	sum := nums[0]
	temp := nums[0]
	for i := 1; i < length; i++ {
		num := nums[i]
		temp = max(num, (num + temp))
		sum = max(sum, temp)
	}
	return sum
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}
