/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 10-11-2019
 * |
 * | File Name:     p6.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println(bulbSwitch(n))
}

func bulbSwitch(n int) int {
	r := math.Sqrt(float64(n))
	return int(math.Floor(r))
}
