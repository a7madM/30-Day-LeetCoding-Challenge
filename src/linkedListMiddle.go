package main

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3290/

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func main() {
	head := listNode{1, nil}
	temp := &head
	for i := 2; i <= 10; i++ {
		node := listNode{i, nil}
		temp.Next = &node
		temp = &node
	}
	middle := middleNode(&head)

	fmt.Println(middle.Val)
}

func middleNode(head *listNode) *listNode {
	length := countLength(head)

	mid := (length / 2) + 1

	midPointer := *head
	for i := 1; i < mid; i++ {
		midPointer = *midPointer.Next
	}

	return &midPointer
}

func countLength(head *listNode) int {
	temp := *head
	length := 0
	for {
		length++
		if temp.Next == nil {
			break
		}
		temp = *temp.Next
	}
	return length
}
