package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}

	points := make([][]int, 0)
	for i := 0; i < n; i++ {
		var x, y int
		if _, err := fmt.Scanf("%d %d", &x, &y); err != nil {
			return
		}
		points = append(points, []int{x, y})
	}
	fmt.Println(maxPoints(points))
}

func maxPoints(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	if len(points) == 0 {
		return 0
	}
	max := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			x1 := points[i][0]
			y1 := points[i][1]
			x2 := points[j][0]
			y2 := points[j][1]

			f := func(x, y int) bool {
				if x2 == x1 && y2 == y1 {
					return x1 == x && y1 == y
				}
				return (x2-x1)*(y-y1) == (y2-y1)*(x-x1)
			}
			c := count(f, points)
			if c > max {
				max = c
			}
		}
	}
	return max
}

func count(f func(x, y int) bool, points [][]int) int {
	c := 0
	for _, p := range points {
		if f(p[0], p[1]) {
			c++
		}
	}
	return c
}
