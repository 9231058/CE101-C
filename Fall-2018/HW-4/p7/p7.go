/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 13-11-2018
 * |
 * | File Name:     750A.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n, k int

	fmt.Scanf("%d %d", &n, &k)

	contestTime := k
	start := 0
	end := n

	for end >= start {
		i := (start + end) / 2
		if 15*i*(i+1)/2 > contestTime {
			end = i - 1
		} else {
			start = i + 1
		}
	}

	fmt.Println(end)
}
