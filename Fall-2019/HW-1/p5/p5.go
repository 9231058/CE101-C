/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-10-2019
 * |
 * | File Name:     p5.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	for i := 10; i < 100; i++ {
		for j := i; j < 100; j++ {
			if isPalindrome(i * j) {
				fmt.Println(j * i)
			}
		}
	}
}

func isPalindrome(n int) bool {
	r := n
	m := 0
	for n != 0 {
		m = m*10 + n%10
		n = n / 10
	}
	return r == m
}
