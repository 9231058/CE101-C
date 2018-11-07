/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 07-11-2018
 * |
 * | File Name:     p3.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var x1, x2, x3, x4 int

	fmt.Scanf("%d %d %d %d", &x1, &x2, &x3, &x4)

	// sorts the numbers by merge sort with descending order
	// nlogn = 4 * 2 = 8
	if x1 < x2 { // sorts x1 and x2
		x1, x2 = x2, x1
	}
	if x3 < x4 { // sorts x3 and x4
		x3, x4 = x4, x3
	}
	// merges
	if x1 > x3 {
		if x2 < x3 {
			if x2 < x4 { // x1 x3 x4 x2
				x3, x2 = x2, x3
				x4, x3 = x3, x4
			} else { // x1 x3 x2 x4
				x3, x2 = x2, x3
			}
		} else { // x1 x2 x3 x4
		}
	} else {
		if x1 < x4 { // x3 x4 x1 x2
			x1, x3 = x3, x1
			x2, x4 = x4, x2
		} else {
			if x2 < x4 { // x3 x1 x4 x2
				x1, x3 = x3, x1
				x2, x3 = x3, x2
				x3, x4 = x4, x3
			} else { // x3 x1 x2 x4
				x1, x3 = x3, x1
				x2, x3 = x3, x2
			}
		}
	}

	fmt.Printf("%d %d %d %d\n", x1, x2, x3, x4)
}
