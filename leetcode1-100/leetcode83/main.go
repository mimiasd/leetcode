package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    temp := head
    for temp != nil {
        if temp.Next != nil {
            if temp.Val == temp.Next.Val {
                temp.Next = temp.Next.Next
            } else {
                temp = temp.Next
            }
        } else {
            break
        }
    }

    return head
}
