package main

import (
	"fmt"
	"math"
	"strconv"
)

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3284/

func main() {
	happy := isHappy(2)
	fmt.Println(happy)
}
func numberDigits(n int) []int {
	digitsStr := strconv.Itoa(n)
	length := len(digitsStr)
	digits := make([]int, length)

	for i := 0; i < length; i++ {
		digit, _ := strconv.Atoi(string(digitsStr[i]))
		digits[i] = digit
	}

	return digits
}

func isHappy(n int) bool {
	fmt.Println(n)
	if n == 4 {
		return false
	}

	digits := numberDigits(n)
	length := len(digits)
	var sum float64
	for i := 0; i < length; i++ {
		sum += math.Pow(float64(digits[i]), 2)
	}
	if sum == 1 {
		return true
	}

	return isHappy(int(sum))
}
