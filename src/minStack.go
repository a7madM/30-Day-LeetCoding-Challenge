package main

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3292/

import "fmt"

func main() {
	obj := Constructor();
	obj.Push(1);
	obj.Push(2);
	obj.Push(3);

	obj.Pop();
	param_3 := obj.Top();
	param_4 := obj.GetMin();

	fmt.Println(param_3, param_4)
}

type MinStack struct {
	stack []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    var minStack = new(MinStack)
    minStack.stack = []int{}
    minStack.minStack = []int{}
    return *minStack
}
func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, x)
    if(len(this.minStack) == 0) {
        this.minStack = append(this.minStack, x)
    } else if x <= this.minStack[len(this.minStack) - 1] {
        this.minStack = append(this.minStack, x)
    }
}

func (this *MinStack) Pop()  {
    if this.stack[len(this.stack) - 1] == this.minStack[len(this.minStack) - 1] {
        this.minStack = this.minStack[:len(this.minStack) - 1]
    }
    this.stack = this.stack[:len(this.stack) - 1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack) - 1]
}

// return the top value from minStack
func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack) - 1]
}

