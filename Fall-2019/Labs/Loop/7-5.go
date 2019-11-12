/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 12-11-2019
 * |
 * | File Name:     7-5.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	sum := 0

	for i := 0; i < n; i++ {
		var c int
		fmt.Scanf("%d", &c)
		sum += c
	}

	fmt.Println((n*(n+1))/2 - sum)
}
