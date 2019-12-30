package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var r *ListNode
	var pre *ListNode

	for head != nil && head.Next != nil {
		if r == nil {
			r = head.Next
		}
		if pre != nil {
			pre.Next = head.Next
		}
		head.Next, head.Next.Next = head.Next.Next, head
		pre = head
		head = head.Next
	}
	return r
}
