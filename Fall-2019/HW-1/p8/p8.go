/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-10-2019
 * |
 * | File Name:     p8.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}

	// answer is C(2n, n) and they can calculate it using its direct formula
	fmt.Println(c(2*n, n))
}

func c(n, m int) int {
	if m == 0 {
		return 1
	}
	if m == 1 {
		return n
	}
	if n == m {
		return 1
	}

	return c(n-1, m-1) + c(n-1, m)
}
