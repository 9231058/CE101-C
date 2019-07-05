/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 21-12-2018
 * |
 * | File Name:     p4.go
 * +===============================================
 */

package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {
	points := make([]point, 10)
	newPoints := make([]point, 0)

	// reads points from stdin
	for i := 0; i < 10; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		points[i] = point{x, y}
	}

	for i, p := range points {
		var isAdded = true // is this point going to be added or not?
		// checks for duplicate x
		for di, dp := range points {
			if p.x == dp.x && di != i {
				if p.y > dp.y {
					isAdded = false
				}
			}
		}
		if isAdded {
			newPoints = append(newPoints, p)
		}
	}

	for _, p := range newPoints {
		fmt.Printf("%d %d\n", p.x, p.y)
	}
}
