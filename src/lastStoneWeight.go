package main

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3297/

import (
	"fmt"
	"sort"
)

func main() {
	stones := []int{7, 1, 5, 3, 6, 4}
	result := lastStoneWeight(stones)

	fmt.Println(result)
}

func lastStoneWeight(stones []int) int {
	// firstly, sort slice of stones
	sort.Ints(stones)

	// implement "Sorted Append" function, which will
	// add element to the slice keeping it sorted
	// using SearchInts func from the stdlib
	// https://golang.org/pkg/sort/#SearchInts
	sortedAppend := func(x int, slice []int) []int {
		// find position where our new element supposed to be located
		i := sort.SearchInts(slice, x)
		// append the new element to the desired position
		//  [0...i, x, i...N]
		return append(slice[:i], append([]int{x}, slice[i:]...)...)
	}

	// crash stones until only one (or none) left
	for l := len(stones); l > 1; l = len(stones) {
		x, y := l-1, l-2           // take two biggest stones
		s := stones[x] - stones[y] // crash them and see what left
		stones = stones[:y]        // remove crashed stones from initial slice
		if s > 0 {                 // if something left from the crashed stones
			// put it in the original slice according to its size
			stones = sortedAppend(s, stones)
		}
	}

	// no stones left
	if len(stones) == 0 {
		return 0
	}

	// return the leftover stone
	return stones[0]
}
