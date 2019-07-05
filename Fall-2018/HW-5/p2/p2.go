/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 23-11-2018
 * |
 * | File Name:     p2.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var x1, y1 int
	var x2, y2 int
	var x3, y3 int
	var x4, y4 int

	fmt.Scanf("%d %d", &x1, &y1)
	fmt.Scanf("%d %d", &x2, &y2)
	fmt.Scanf("%d %d", &x3, &y3)
	fmt.Scanf("%d %d", &x4, &y4)

	// Circular Permutation over free circle is (n - 1)!/2
	if rt(x1, y1, x2, y2, x4, y4) && rt(x2, y2, x1, y1, x3, y3) && rt(x3, y3, x2, y2, x4, y4) && rt(x4, y4, x3, y3, x1, y1) {
		fmt.Println("True")
		// p2 p3
		// p1 p4
	} else if rt(x1, y1, x3, y3, x4, y4) && rt(x2, y2, x4, y4, x3, y3) && rt(x3, y3, x2, y2, x1, y1) && rt(x4, y4, x2, y2, x1, y1) {
		fmt.Println("True")
		// p4 p2
		// p1 p3
	} else if rt(x1, y1, x3, y3, x2, y2) && rt(x2, y2, x4, y4, x1, y1) && rt(x3, y3, x4, y4, x1, y1) && rt(x4, y4, x2, y2, x3, y3) {
		fmt.Println("True")
		// p2 p4
		// p1 p3
	} else {
		fmt.Println("False")
	}
}

func rt(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int) bool {
	a := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	b := (x1-x3)*(x1-x3) + (y1-y3)*(y1-y3)
	c := (x2-x3)*(x2-x3) + (y2-y3)*(y2-y3)

	return a+b == c
}
