package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    var tempHead *ListNode = nil
    cur := head

    for cur != nil {
        t := cur.Next
        cur.Next = tempHead
        tempHead = cur
        cur = t
    }

    return tempHead
}
