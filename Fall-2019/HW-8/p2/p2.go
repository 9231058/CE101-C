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

	l2 := swapPairs(l1)
	for l2 != nil {
		fmt.Printf("%d ", l2.Val)
		l2 = l2.Next
	}
	fmt.Println()
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
