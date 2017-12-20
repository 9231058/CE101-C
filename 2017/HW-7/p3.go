/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 20-12-2017
 * |
 * | File Name:     p3.go
 * +===============================================
 */

package main

import (
	"fmt"
	"sort"
)

// Quotient represents fraction of two integer
type Quotient struct {
	Numerator   int
	Denominator int
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	q := make([]Quotient, n)

	for i := 0; i < n; i++ {
		var num, denom int
		fmt.Scanf("%d %d", &num, &denom)

		q[i] = Quotient{num, denom}
	}

	var o int
	fmt.Scanf("%d", &o)

	sort.Slice(q, func(i, j int) bool {
		if q[i].Numerator*q[j].Denominator < q[j].Numerator*q[j].Denominator {
			if o == 1 {
				return true
			}
			return false
		}
		if o == 1 {
			return false
		}
		return true
	})

	for _, q := range q {
		fmt.Printf("%d/%d\n", q.Numerator, q.Denominator)
	}
}
