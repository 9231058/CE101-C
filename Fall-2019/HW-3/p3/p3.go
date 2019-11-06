/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 06-11-2019
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
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}
	fmt.Println(isPowerTwo(n))
}

func isPowerTwo(n int) bool {
	if n == 0 {
		return false
	}
	l := math.Log2(float64(n))
	return l == math.Round(l)
}
