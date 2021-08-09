package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	 n := len(lists)
    if n == 0 {
        return nil
    }

	for i := 1; i < n; i++ {
		lists[i] = mergeTwoLists(lists[i-1], lists[i])
	}

	return lists[n-1]
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	dummy := &ListNode {
		Val: 0,
		Next: nil,
	}

	temp := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}	
		temp = temp.Next
	} 
	if l1 == nil {
		temp.Next = l2
	} else {
		temp.Next = l1
	}

	return dummy.Next
}
