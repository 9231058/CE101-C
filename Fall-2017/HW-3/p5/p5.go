/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 05-11-2017
 * |
 * | File Name:     p5.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, w1, h1 float64
	var x2, y2, w2, h2 float64

	var w, h float64

	fmt.Scanf("%f %f %f %f", &x1, &y1, &w1, &h1)
	fmt.Scanf("%f %f %f %f", &x2, &y2, &w2, &h2)

	if x1 < x2 {
		w = math.Max(math.Min(x1+w1-x2, w2), 0)
	} else {
		w = math.Max(math.Min(x2+w2-x1, w1), 0)
	}

	if y1 < y2 {
		h = math.Max(math.Min(y1+h1-y2, h2), 0)
	} else {
		h = math.Max(math.Min(y2+h2-y1, h1), 0)
	}

	fmt.Printf("%.2f\n", h*w)
}
