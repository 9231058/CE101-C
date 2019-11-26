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

func largestFactor(n int) int {
	for i := n; i >= 1; i-- {
		if prime(i) == true && n%i == 0 {
			return i
		}
	}
	return 1
}
func main() {
	var a int
	if _, err := fmt.Scanf("%d\n", &a); err != nil {
		return
	}

	for a > 1 {
		f := largestFactor(a)
		fmt.Println(f)
		a /= f
	}
}
