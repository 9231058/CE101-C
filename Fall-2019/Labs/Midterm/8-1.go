/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 21-11-2019
 * |
 * | File Name:     8-1.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(distance(0, 0, 1, 1))
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(
		float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)),
	)
}
