/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 11-11-2019
 * |
 * | File Name:     8-6.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var p int

	fmt.Scanf("%d", &p)

	sp := math.Sqrt(float64(p))
	fmt.Println(sp == math.Round(sp))
}
