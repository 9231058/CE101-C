/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 21-11-2019
 * |
 * | File Name:     p1.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var x1, x2, x3, x4 int

	fmt.Scanf("%d %d %d %d\n", &x1, &x2, &x3, &x4)

	max := x1
	if x2 > max {
		max = x2
	}
	if x3 > max {
		max = x3
	}
	if x4 > max {
		max = x4
	}

	if max-x1 != 0 {
		fmt.Printf("%d ", max-x1)
	}
	if max-x2 != 0 {
		fmt.Printf("%d ", max-x2)
	}
	if max-x3 != 0 {
		fmt.Printf("%d ", max-x3)
	}
	if max-x4 != 0 {
		fmt.Printf("%d ", max-x4)
	}
	fmt.Println()
}
