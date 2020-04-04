package main

import "fmt"

func main() {
	fmt.Println("Hello")
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	length := len(nums)
	count := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}

	for i := count; i < length; i++ {
		nums[i] = 0
	}
}
