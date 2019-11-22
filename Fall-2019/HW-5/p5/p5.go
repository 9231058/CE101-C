/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 22-11-2019
 * |
 * | File Name:     p5.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}
	fmt.Println(binaryGap(n))
}

func binaryGap(n int) int {
	latest := -1
	longest := 0

	for n > 0 {
		current := largestPower2(&n)
		if latest == -1 {
			latest = current
		} else {
			if longest < latest-current {
				longest = latest - current
			}
			latest = current
		}
	}
	return longest
}

func largestPower2(n *int) int {
	i := 1
	c := 0
	for i <= *n {
		i = i * 2
		c++
	}
	*n -= i / 2
	return c - 1
}
