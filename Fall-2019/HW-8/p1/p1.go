package main

import "fmt"

// ListNode is a linked-list node
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1, p1 *ListNode
	var l2, p2 *ListNode

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

	for {
		var n int
		if _, err := fmt.Scanf("%d", &n); err != nil {
			return
		}
		if n == 0 {
			break
		}
		if l2 == nil {
			p2 = &ListNode{
				Val:  n,
				Next: nil,
			}
			l2 = p2
		} else {
			p2.Next = &ListNode{
				Val:  n,
				Next: nil,
			}
			p2 = p2.Next
		}
	}

	l3 := mergeTwoLists(l1, l2)
	for l3 != nil {
		fmt.Printf("%d ", l3.Val)
		l3 = l3.Next
	}
	fmt.Println()
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
