/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 23-11-2018
 * |
 * | File Name:     p1.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var a1, b1, c1 float64
	var a2, b2, c2 float64

	fmt.Scanf("%f %f %f", &a1, &b1, &c1)
	fmt.Scanf("%f %f %f", &a2, &b2, &c2)

	determinant := a1*b2 - a2*b1
	if determinant != 0 {
		x1 := (c1*b2 - c2*b1) / determinant
		x2 := (c2*a1 - c1*a2) / determinant

		fmt.Printf("%.02f %.02f\n", x1, x2)
	} else {
		fmt.Println("No Solution")
	}
}
