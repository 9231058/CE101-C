package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	var current = 0
	var max = 0
	var previous = 0

	for i := 0; i < n; i++ {
		var next int
		fmt.Scanf("%d", &next)
		if next >= previous {
			current++
		} else {
			if current > max {
				max = current
			}
			current = 1
		}
		previous = next
	}
	if current > max {
		max = current
	}

	fmt.Printf("%d\n", max)
}