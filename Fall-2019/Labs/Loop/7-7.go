/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 12-11-2019
 * |
 * | File Name:     7-7.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	fmt.Println(sum == n)
}
