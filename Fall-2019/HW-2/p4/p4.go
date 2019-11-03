/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 04-11-2019
 * |
 * | File Name:     p4.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int

	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}

	n = n / 2

	// this loop generates four even random number in range [m, n]
	for i := 0; i < 4; i++ {
		r := rand.Intn(n + 1) // an integer between 0 to n/2
		fmt.Println(2 * r)    // [0, n/2] and 2 * r is always even
	}
}
