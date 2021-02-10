/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 05-12-2020
 * |
 * | File Name:     p3.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64

	var ia, ib int
	if _, err := fmt.Scanf("%d %d %f %f", &ia, &ib, &c, &d); err != nil {
		return
	}

	a = float64(ia)
	b = float64(ib)

	if c > d {
		fmt.Println("Theta is bigger than theta max!")
		return
	}

	result := math.Sqrt(a*a + b*b - (2 * a * b * math.Cos(c)))

	fmt.Printf("%.2f\n", result)
}
