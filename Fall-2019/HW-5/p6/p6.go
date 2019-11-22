/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 22-11-2019
 * |
 * | File Name:     p6.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}
	fmt.Println(primePalindrome(n))
}

func primePalindrome(N int) int {
	i := N
	for {
		if isPalindrome(i) && isPrime(i) {
			return i
		}
		i++
	}
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPalindrome(n int) bool {
	b := n
	r := 0

	for n > 0 {
		d := n % 10
		r = r*10 + d
		n = n / 10
	}

	return r == b
}
