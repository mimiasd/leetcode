package main

type MyQueue struct {
	inStack []int
	outStack []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		inStack: make([]int, 0),
		outStack: make([]int, 0),
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) in2Out() {
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		this.in2Out()
	}

	x := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return x
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.in2Out()
	}

	return this.outStack[len(this.outStack)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

