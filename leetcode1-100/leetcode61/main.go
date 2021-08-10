package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	size := 0
	cur := head

	for ; cur.Next != nil; cur = cur.Next {
		size++
	}

	size++
	k = k % size
	cur.Next = head
	cur = head
	for i := 1; i < size - k; i++ {
		cur = cur.Next
	}
	res := cur.Next
	cur.Next = nil

	return res
}
