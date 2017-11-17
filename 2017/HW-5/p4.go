/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 17-11-2017
 * |
 * | File Name:     p4.go
 * +===============================================
 */

package main

import "fmt"

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func nextPrime(n int) int {
	if n == 2 {
		return 3
	}
	n += 2
	for !isPrime(n) {
		n += 2
	}
	return n
}

func main() {
	var n int

	fmt.Scanf("%d", &n)

	for i := 4; i <= n; i += 2 {
		fmt.Printf("Even number: %d\n", i)
		for j := 2; j <= i/2; {
			if isPrime(i - j) {
				fmt.Printf("Prime 1: %d\n", j)
				fmt.Printf("Prime 2: %d\n", i-j)
				break
			}
			j = nextPrime(j)
		}
	}
}
