/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 11-11-2017
 * |
 * | File Name:     p2.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	p := make([]int, 0)

	fmt.Scanf("%d", &n)

	for i := 2; i <= n; i++ {
		if isPerfect(i) {
			p = append(p, i)
		}
	}
	fmt.Println(len(p))
	for _, i := range p {
		fmt.Println(i)
	}
}

func isPerfect(n int) bool {
	s := 1
	for i := 2; i < n; i++ {
		if n%i == 0 {
			s += i
		}
	}
	if s == n {
		return true
	}
	return false
}
