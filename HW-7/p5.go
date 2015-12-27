package main

import "fmt"

func gcd(a int, b int) int {
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
}

type Fraction struct {
	Numerator   int
	Denominator int
}

func (fs Fraction) Plus(fe Fraction) Fraction {
	var sum Fraction
	sum.Numerator = (fs.Numerator * fe.Denominator) + (fs.Denominator * fe.Numerator)
	sum.Denominator = fs.Denominator * fe.Denominator
	gcd := gcd(sum.Numerator, sum.Denominator)
	sum.Numerator /= gcd
	sum.Denominator /= gcd
	return sum
}

func (f Fraction) String() string {
	return fmt.Sprintf("%d / %d", f.Numerator, f.Denominator)
}

func (fs Fraction) LessWithScale(fe Fraction, x float64) bool {
	return float64(fs.Numerator*fe.Denominator) < x*float64(fe.Numerator*fs.Denominator)
}

type Node struct {
	Info interface{}
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *List) Append(v interface{}) {
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

func (l *List) DeleteValue(v interface{}) {
	if l.Head.Info == v {
		l.Head = l.Head.Next
		l.Size--
		return
	}

	var p *Node = l.Head
	var c *Node = p.Next

	for c != nil {
		if c.Info == v {
			p.Next = c.Next
			l.Size--
			return
		}
		p = p.Next
		c = c.Next
	}
}

func (l *List) Print() {
	for c := l.Head; c != nil; c = c.Next {
		fmt.Printf("%v ", c.Info)
	}
	fmt.Printf("\n")
}

func main() {
	var list List
	sum := Fraction{
		Denominator: 1,
		Numerator:   0,
	}
	for {
		var f Fraction
		fmt.Scanf("%d %d", &f.Numerator, &f.Denominator)
		if f.Numerator == 0 && f.Denominator == 0 {
			break
		}
		sum = sum.Plus(f)
		list.Append(f)
	}
	var x float64
	fmt.Scanf("%f", &x)

	for c := list.Head; c != nil; c = c.Next {
		if c.Info.(Fraction).LessWithScale(sum, x) {
			fmt.Println(c.Info)
		}
	}
}
