package main

type MyStack struct {
	queue []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack {
		queue: make([]int, 0),
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.queue = append(this.queue, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	x := this.queue[len(this.queue)-1]
	this.queue = this.queue[:len(this.queue)-1]

	return x
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1] 
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
