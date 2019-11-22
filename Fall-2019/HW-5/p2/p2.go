/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 22-11-2019
 * |
 * | File Name:     p2.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}
	fmt.Println(climbStairs(n))
}

var mem = make(map[int]int)

func climbStairs(n int) int {
	r, ok := mem[n]
	if ok {
		return r
	}

	if n == 1 || n == 0 {
		return 1
	}

	r = climbStairs(n-1) + climbStairs(n-2)
	mem[n] = r
	return r
}
