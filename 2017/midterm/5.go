/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 03-12-2017
 * |
 * | File Name:     5.go
 * +===============================================
 */

package main

import "fmt"

func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibFactor(n int) bool {
	if n == 0 {
		return true
	}
	if n == 1 {
		fmt.Println(1)
		return true
	}

	i := 2

	for fib(i) <= n {
		if fibFactor(n-fib(i)) == true {
			fmt.Println(i)
			return true
		}
		i++
	}
	return false
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	if fibFactor(n) == false {
		fmt.Println("No fibonacci factorization was found")
	}
}
