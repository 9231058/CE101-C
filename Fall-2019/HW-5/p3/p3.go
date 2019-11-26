/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 26-11-2019
 * |
 * | File Name:     p3.go
 * +===============================================
 */

package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}

	p := point{
		x: 2,
		y: 1,
	}
	for i := 0; i < n; i++ {
		p = transform(p)
	}
	fmt.Printf("(%d, %d)\n", p.x, p.y)
}

func transform(p point) point {
	return point{
		x: p.x*p.x - p.y*p.y,
		y: 2 * p.x * p.y,
	}
}
