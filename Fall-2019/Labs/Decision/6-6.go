/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 05-11-2019
 * |
 * | File Name:     6-6.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}

	fmt.Println(addDigits(n))
}

func addDigits(num int) int {
	r := num % 9
	if r == 0 {
		r = 9
	}
	if num == 0 {
		r = 0
	}
	return r
}
