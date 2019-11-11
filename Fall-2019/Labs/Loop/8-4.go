/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 11-11-2019
 * |
 * | File Name:     8-4.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n, m int

	if _, err := fmt.Scanf("%d %d", &n, &m); err != nil {
		return
	}

	// k = math.Pow(m, 1/n)
	start := 0
	end := m
	for end-start > 1 {
		k := (start + end) / 2

		if pow(k, n) <= m {
			start = k
		} else {
			end = k
		}
	}
	fmt.Println(start)
}

func pow(k, n int) int {
	r := 1
	for i := 0; i < n; i++ {
		r = r * k
	}
	return r
}
