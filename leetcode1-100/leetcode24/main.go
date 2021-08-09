package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode {
		Val: -1,
		Next: head,
	}

	cur := head
	pre := dummy

	for cur != nil && cur.Next != nil {
		t := cur.Next
		cur.Next = t.Next
		t.Next = cur
		pre.Next = t
		pre = cur
		cur = cur.Next
			
	}
	return dummy.Next
}
