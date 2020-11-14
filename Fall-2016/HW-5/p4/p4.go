/*
 * +===============================================
 * | Author:        Elahe Jalalpour (el.jalalpour@gmail.com)
 * |
 * | Creation Date: 27-11-2015
 * |
 * | File Name:     p4.go
 * +===============================================
 */
package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d\n", &n)

	x := 2
	y := 1

	for i := 0; i < n; i++ {
		x, y = coordinate(x, y)
	}

	fmt.Printf("(%d, %d)\n", x, y)
}

func coordinate(x int, y int) (int, int) {
	var u = x*x - y*y
	var v = 2 * x * y
	return u, v
}
