/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 26-11-2019
 * |
 * | File Name:     p7.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}

	fmt.Println(f1(n))
	fmt.Println(f2(n))
	fmt.Println(f3(n))
}

func f1(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return f1(n-1) + f1(n-2)
}

func f2(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return f2(n-2) + f2(n-3) - f2(n-1)
}

func f3(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return 2*f3(n-2) - f3(n-1) + f3(n-3)
}
