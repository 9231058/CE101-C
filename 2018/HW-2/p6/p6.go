/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 26-10-2018
 * |
 * | File Name:     p6.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var m, n int

	fmt.Scanf("%d %d", &n, &m)

	if m >= 0 || n <= 0 {
		panic("0 < n < RAND_MAX and -RAND_MAX < m < 0")
	}

	if n%2 == 0 {
		n = n / 2
	} else {
		n = (n - 1) / 2
	}

	if m%2 == 0 {
		m = m / 2
	} else {
		m = (m - 1) / 2
	}

	// this loop generates four even random number in range [m, n]
	for i := 0; i < 4; i++ {
		// by n and m in this comment we mean they original value
		// [0, n/2 + (-m/2)] -> [m/2, n/2]
		r := rand.Intn(n+(-m)+1) + m
		fmt.Println(2 * r) // [m, n] and 2 * r is always even
	}
}
