/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 17-11-2017
 * |
 * | File Name:     p3.go
 * +===============================================
 */

package main

import "fmt"

var combinationCalls int

func combination(n int, m int) int {
	combinationCalls++
	if n == 0 && m == 0 {
		return 1
	}
	if n == 0 {
		return 0
	}
	if m == 0 {
		return 1
	}
	return combination(n-1, m) + combination(n-1, m-1)
}

func main() {
	var n, m int

	fmt.Scanf("%d %d", &n, &m)

	fmt.Println(combination(n, m))
	fmt.Println(combinationCalls)
}
