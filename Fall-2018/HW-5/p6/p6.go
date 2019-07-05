/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 21-11-2018
 * |
 * | File Name:     996A.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(bill(n))
}

func bill(amount int) int {
	var bill50, bill10, bill5, bill2, bill1 int
	var min = math.MaxInt32

	if amount >= 50 {
		bill50 = bill(amount-50) + 1
		if min > bill50 {
			min = bill50
		}
	}
	if amount >= 10 {
		bill10 = bill(amount-10) + 1
		if min > bill10 {
			min = bill10
		}

	}
	if amount >= 5 {
		bill5 = bill(amount-5) + 1
		if min > bill5 {
			min = bill5
		}

	}
	if amount >= 2 {
		bill2 = bill(amount-2) + 1
		if min > bill2 {
			min = bill2
		}

	}
	if amount >= 1 {
		bill1 = bill(amount-1) + 1
		if min > bill1 {
			min = bill1
		}
	}
	if amount == 0 {
		return 0
	}
	return min
}
