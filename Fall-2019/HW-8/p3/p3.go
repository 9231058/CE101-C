package main

import "fmt"

// ListNode is a linked-list node
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1, p1 *ListNode

	for {
		var n int
		if _, err := fmt.Scanf("%d", &n); err != nil {
			return
		}
		if n == 0 {
			break
		}
		if l1 == nil {
			p1 = &ListNode{
				Val:  n,
				Next: nil,
			}
			l1 = p1
		} else {
			p1.Next = &ListNode{
				Val:  n,
				Next: nil,
			}
			p1 = p1.Next
		}
	}

	fmt.Println(isPalindrome(l1))
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
