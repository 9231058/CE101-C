package main

import "fmt"

type Node struct {
	info int
	next *Node
}

func Create_Node(i int) *Node {
	var nn Node
	nn.info = i
	nn.next = nil
	return &nn
}

func Del_Node(a int, list *Node) {
	var parent *Node = list
	var currentlist *Node = parent.next

	if parent.info == a {
		list = currentlist
	}
	for currentlist.info != a {
		parent = parent.next
		currentlist = parent.next
	}
	parent.next = currentlist.next

}

func main() {
	var list *Node = Create_Node(0)
	list.next = Create_Node(10)
	list.next.next = Create_Node(20)
	list.next.next.next = Create_Node(30)
	list.next.next.next.next = Create_Node(40)
	list.next.next.next.next.next = Create_Node(50)

	Del_Node(10, list)

	for i := list; i.next != nil; i = i.next {
		fmt.Println(i.info)
	}

}
