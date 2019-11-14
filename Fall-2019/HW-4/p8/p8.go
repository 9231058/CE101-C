/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 14-11-2019
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
	fmt.Println(countDigitOne(n))
}

func countDigitOne(n int) int {
	p10 := 1
	r := 0
	for p10 <= n {
		p10 *= 10
		r += (n / p10) * (p10 / 10)
		c := n%p10 - p10/10

		if c >= 0 {
			if c >= p10/10 {
				r += (p10 / 10)
			} else {
				r += c + 1
			}
		}
	}
	return r
}
