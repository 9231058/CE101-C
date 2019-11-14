/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 14-11-2019
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
	var c int

	if _, err := fmt.Scanf("%d", &c); err != nil {
		return
	}
	fmt.Println(judgeSquareSum(c))
}

func judgeSquareSum(c int) bool {
	for i := 0; i*i <= c; i++ {
		r := math.Sqrt(float64(c - i*i))
		if r == math.Round(r) {
			return true
		}
	}
	return false
}
