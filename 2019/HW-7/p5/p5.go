package main

import "fmt"

func main() {
	count := make(map[int]int)

	for {
		var x int
		fmt.Scanf("%d", &x)
		if x == 0 {
			break
		}
		count[x]++
	}

	for x, n := range count {
		fmt.Printf("%d %d\n", n, x)
	}
}
