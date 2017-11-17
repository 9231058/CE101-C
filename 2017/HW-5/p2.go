/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 17-11-2017
 * |
 * | File Name:     p2.go
 * +===============================================
 */

package main

import (
	"fmt"
	"sort"
)

func prime(p int) bool {
	for i := 2; i < p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func maxFactor(n int) int {
	for i := n; i >= 2; i-- {
		if n%i == 0 && prime(i) {
			return i
		}
	}
	return 0
}

func allFactor(n int) map[int]int {
	if n == 1 {
		return make(map[int]int)
	}
	l := maxFactor(n)
	m := allFactor(n / l)
	m[l]++
	return m
}

func main() {
	var m int
	fmt.Scanf("%d", &m)
	for i := 1; i <= m; i++ {
		fmt.Printf("%d = 1", i)

		m := allFactor(i)
		var ks []int
		for k := range m {
			ks = append(ks, k)
		}
		sort.Ints(ks)

		for _, k := range ks {
			fmt.Printf(" * %d^%d", k, m[k])
		}
		fmt.Println()
	}
}
