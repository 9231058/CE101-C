package main

import "fmt"

/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 21-11-2019
 * |
 * | File Name:     8-6.go
 * +===============================================
 */

func main() {
	var n int
	var dir int

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &dir)

	current := 0
	longest := 0
	latest := 0
	for i := 0; i < n; i++ {
		var e int
		fmt.Scanf("%d", &e)
		if i == 0 {
			latest = e
			current = 1
			continue
		}
		if dir == 1 {
			if e >= latest {
				current++
			} else {
				current = 1
			}
		}
		if dir == -1 {
			if e <= latest {
				current++
			} else {
				current = 1
			}
		}
		if current > longest {
			longest = current
		}
		latest = e
	}
	fmt.Println(longest)
}
