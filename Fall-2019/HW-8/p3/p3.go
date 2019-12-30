package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var pre *ListNode
	for head != nil {
		it := head

		for it.Next != nil {
			pre = it
			it = it.Next
		}

		if it.Val != head.Val {
			return false
		}

		if pre != nil {
			pre.Next = nil
		}
		head = head.Next
	}
	return true
}
