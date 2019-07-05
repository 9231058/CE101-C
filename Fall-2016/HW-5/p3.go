/*
* +===============================================
* | Author:        Parham Alvani (parham.alvani@gmail.com)
* |
* | Creation Date: 09-01-2017
* |
* | File Name:     p3.go
* +===============================================
 */
package main

import (
	"fmt"
)

type Hanoi struct {
	L [3][]int
}

type HanoiSolver interface {
	Solve(start int, end, int, size int)
	Move(start int, end int)
	Print()
}

func (h *Hanoi) Print() {
	ring := 0
	l := 0
	for i := 0; i < len(h.L); i++ {
		ring += len(h.L[i])
		if l < len(h.L[i]) {
			l = len(h.L[i])
		}
	}
	l--
	for ; l >= 0; l-- {
		for i := 0; i < len(h.L); i++ {
			if l < len(h.L[i]) {
				for j := 0; j < h.L[i][l]; j++ {
					fmt.Printf("*")
				}
				for j := 0; j < ring-h.L[i][l]; j++ {
					fmt.Printf(" ")
				}
				fmt.Printf("   ")
			} else {
				for j := 0; j < ring; j++ {
					fmt.Printf(" ")
				}
				fmt.Printf("   ")
			}
		}
		fmt.Printf("\n")
	}
	for i := 0; i < len(h.L); i++ {
		fmt.Printf("%c", byte(i+'A'))
		for j := 0; j < ring-1; j++ {
			fmt.Printf(" ")
		}
		fmt.Printf("   ")
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
}

func (h *Hanoi) Solve(start int, end int, size int) {
	if size == 1 {
		h.Move(start, end)
	} else {
		h.Solve(start, 3-(end+start), size-1)
		h.Move(start, end)
		h.Solve(3-(start+end), end, size-1)
	}
}

func (h *Hanoi) Move(start int, end int) {
	ring := h.L[start][len(h.L[start])-1]
	h.L[start] = h.L[start][:len(h.L[start])-1]
	h.L[end] = append(h.L[end], ring)
	h.Print()
}

func main() {
	var N int
	fmt.Scanf("%d", &N)

	h := new(Hanoi)
	for i := N; i > 0; i-- {
		h.L[0] = append(h.L[0], i)
	}
	h.Print()
	h.Solve(0, 2, N)
}
