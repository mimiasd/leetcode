package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }

    dummy := &ListNode {
        Val: -1,
        Next: head,
    }
    temp := dummy
    for temp.Next != nil && temp.Next.Next != nil {
        if temp.Next.Val == temp.Next.Next.Val {
            val := temp.Next.Val
            for temp.Next != nil && temp.Next.Val == val {
                temp.Next = temp.Next.Next
            }
        } else {
            temp = temp.Next
        }
    }

    return dummy.Next
}
