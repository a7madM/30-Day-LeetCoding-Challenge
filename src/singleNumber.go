package main

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3283/

func singleNumber(nums []int) int {
	var dictionary map[int]int
	dictionary = make(map[int]int)

	length := len(nums)
	result := 0
	for i := 0; i < length; i++ {
		value := nums[i]
		dictionary[value] = dictionary[value] + 1
	}
	for i := 0; i < length; i++ {
		value := nums[i]

		if dictionary[value] == 1 {
			result = value
		}
	}
	return result
}
