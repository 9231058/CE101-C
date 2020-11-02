/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 27-11-2015
 * |
 * | File Name:     p2-b.go
 * +===============================================
 */
package main

import "fmt"

func prime(n int) bool {
	for j := 2; j < n; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}

func largest_factor(n int) int {
	for i := n; i >= 1; i-- {
		if prime(i) == true && n%i == 0 {
			return i
		}
	}
	return 1
}
func main() {
	var a int
	fmt.Scanf("%d\n", &a)
	for a > 1 {
		f := largest_factor(a)
		fmt.Println(f)
		a /= f
	}
}
