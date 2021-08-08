package main

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	carry := 0
	cur := dummy

	for l1 != nil || l2 != nil || carry != 0 {
		s := carry
		if l1 != nil {
			s += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			s += l2.Val
			l2 = l2.Next
		}

		carry = s / 10
		cur.Next = &ListNode{s % 10, nil}
		cur = cur.Next
	}

	return dummy.Next
}
