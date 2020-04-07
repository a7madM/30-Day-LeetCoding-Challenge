package main

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3289/
import "fmt"

func main() {
	elements := []int{1, 2, 3}
	elements = []int{1, 1, 2}
	elements = []int{1, 1, 2, 2}
	elements = []int{1, 1, 3, 3, 5, 5, 7, 7}
	elements = []int{1, 3, 2, 3, 5, 0}

	count := countElements(elements)

	fmt.Println(count)
}

func countElements(elements []int) int {
	var count int
	length := len(elements)
	var dictionary map[int]int
	dictionary = make(map[int]int)
	for i := 0; i < length; i++ {
		sum := elements[i] + 1
		dictionary[sum] = dictionary[sum] + 1
	}

	for i := 0; i < length; i++ {
		existsInHash, _ := dictionary[elements[i]]
		if existsInHash > 0 {
			dictionary[elements[i]] = 0
			count += existsInHash
		}
	}
	return count
}
