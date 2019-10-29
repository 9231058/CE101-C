/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-10-2019
 * |
 * | File Name:     p9.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	start := 0
	end := 1024

	bottles := make(map[int]bool)

	// the index of poisoned bottle
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}
	bottles[n] = true

	for start+1 < end {
		fmt.Printf("start = %d end = %d\n", start, end)
		mid := (start + end) / 2
		// predecease drinks bottles from start to mid
		for i := start; i < mid; i++ {
			if bottles[i] == true {
				// predecease is dead
				end = mid
			}
		}
		// predecease is alive
		if end != mid {
			start = mid
		}
	}
	fmt.Println(start)
}
