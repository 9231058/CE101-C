package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var r = &ListNode{
		Val:  0,
		Next: nil,
	}
	var l = r

	for l1 != nil && l2 != nil {
		l.Next = &ListNode{
			Val:  0,
			Next: nil,
		}

		l = l.Next
		if l1.Val < l2.Val {
			l.Val = l1.Val
			l1 = l1.Next
		} else {
			l.Val = l2.Val
			l2 = l2.Next
		}
	}
	for l1 != nil {
		l.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
		l = l.Next

		l.Val = l1.Val
		l1 = l1.Next
	}
	for l2 != nil {
		l.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
		l = l.Next

		l.Val = l2.Val
		l2 = l2.Next
	}

	return r.Next
}
