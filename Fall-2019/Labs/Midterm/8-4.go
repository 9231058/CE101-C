/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 21-11-2019
 * |
 * | File Name:     8-4.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	i := 1
	c := 0
	for c < n {
		if isUgly(i) {
			c++
		}
		i++
	}
	fmt.Println(i - 1)
}

func isUgly(n int) bool {
	for n%2 == 0 {
		n /= 2
	}

	for n%3 == 0 {
		n /= 3
	}

	for n%5 == 0 {
		n /= 5
	}

	return n == 1
}
