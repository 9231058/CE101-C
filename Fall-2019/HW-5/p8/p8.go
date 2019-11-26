/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 26-11-2019
 * |
 * | File Name:     p8.go
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

	start := math.MinInt32
	end := math.MaxInt32
	for i := 0; i < n; i++ {
		var x int
		var y int
		if _, err := fmt.Scanf("(%d, %d)\n", &x, &y); err != nil {
			fmt.Println(err)
			return
		}
		if start < x {
			start = x
		}
		if end > y {
			end = y
		}
	}
	fmt.Printf("(%d, %d)\n", start, end)
}
