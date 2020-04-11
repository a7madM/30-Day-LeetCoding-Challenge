package main

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3293/

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{1, nil, nil}
	node := TreeNode{2, nil, nil}
	root.Left = &node
	node = TreeNode{3, nil, nil}
	root.Right = &node
	diameter := diameterOfBinaryTree(&root)
	fmt.Println(diameter)
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var diameter int
	postOrder(root, &diameter)
	return diameter - 1
}

func postOrder(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	left := postOrder(root.Left, diameter)
	right := postOrder(root.Right, diameter)
	*diameter = max(*diameter, left+right+1)
	return max(left, right) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
