/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 27-11-2015
 * |
 * | File Name:     p2-a.go
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
	ans := 1
	var a int
	var b int
	fmt.Scanf("%d %d\n", &a, &b)
	for a != 1 && b != 1 {
		af := largest_factor(a)
		bf := largest_factor(b)
		if af == bf {
			a /= af
			b /= bf
			ans *= af
		} else if af > bf {
			a /= af
		} else {
			b /= bf
		}
	}
	fmt.Println(ans)
}
