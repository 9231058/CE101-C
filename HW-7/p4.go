package main

import "fmt"

type Node struct {
	Info int
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *List) Append(v int) {
	var nn Node
	nn.Info = v
	nn.Next = nil

	if l.Size == 0 {
		l.Tail = &nn
		l.Head = l.Tail
	} else {
		l.Tail.Next = &nn
		l.Tail = &nn
	}
	l.Size++
}

func (l *List) DeleteValue(v int) {
	if l.Head.Info == v {
		l.Head = l.Head.Next
		return
	}

	var p *Node = l.Head
	var c *Node = p.Next

	for c != nil {
		if c.Info == v {
			p.Next = c.Next
			return
		}
		p = p.Next
		c = c.Next
	}
}

func (l *List) Print() {
	for c := l.Head; c != nil; c = c.Next {
		fmt.Printf("%d ", c.Info)
	}
	fmt.Printf("\n")
}

func main() {
	var list List
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Append(50)
	list.Print()

	list.DeleteValue(50)
	list.Print()

	list.DeleteValue(10)
	list.Print()

	list.DeleteValue(40)
	list.Print()

	list.DeleteValue(30)
	list.Print()
}
