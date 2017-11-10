/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 10-11-2017
 * |
 * | File Name:     p1.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var x, y int

	fmt.Scanf("%d %d", &x, &y)

	if x < y {
		x, y = y, x
	}

	for i := 2; i <= x; i++ {
		if x%i == y%i {
			fmt.Printf("%d\n", i)
		}
	}
}
