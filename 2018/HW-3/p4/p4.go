/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 08-11-2018
 * |
 * | File Name:     p4.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1 float64
	var x2, y2 float64
	var x3, y3 float64

	fmt.Scanf("%f %f", &x1, &y1)
	fmt.Scanf("%f %f", &x2, &y2)
	fmt.Scanf("%f %f", &x3, &y3)

	// calculates the square of the imaginary triangle's sides
	s12s := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	s13s := (x1-x3)*(x1-x3) + (y1-y3)*(y1-y3)
	s23s := (x2-x3)*(x2-x3) + (y2-y3)*(y2-y3)

	// is triangle real?
	if math.Sqrt(s12s)+math.Sqrt(s13s) <= math.Sqrt(s23s) ||
		math.Sqrt(s12s)+math.Sqrt(s23s) <= math.Sqrt(s13s) ||
		math.Sqrt(s13s)+math.Sqrt(s23s) <= math.Sqrt(s12s) {
		fmt.Println(0)
		return
	}

	// has two equal sides
	var isosceles bool
	if s12s == s13s || s13s == s23s || s23s == s12s {
		isosceles = true
	}

	// has one angle = 90 degree
	var right bool
	if s12s+s13s == s23s || s12s+s23s == s13s || s23s+s13s == s12s {
		right = true
	}

	if isosceles {
		if right {
			fmt.Println(2)
		} else {
			fmt.Println(1)
		}
	} else {
		fmt.Println(4)
	}
}
