/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 12-11-2019
 * |
 * | File Name:     8-9.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	fmt.Println(countNumbersWithUniqueDigits(n))
}

func countNumbersWithUniqueDigits(n int) int {
	c := 1
	for i := 1; i <= n; i++ {
		p := 9
		for j := 0; j < i-1; j++ {
			p *= (9 - j)
		}
		c += p
	}
	return c
}
