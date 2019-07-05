/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-10-2017
 * |
 * | File Name:     p4.go
 * +===============================================
 */

package main

import (
	"fmt"
)

func main() {
	var weight int
	var cost int

	fmt.Scanf("%d", &weight)

	if weight < 100 {
		cost = 1000
	} else if weight < 200 {
		cost = 1500
	} else if weight < 300 {
		cost = 2000
	} else if weight < 500 {
		cost = 3000
	} else {
		cost = 3000
		cost += ((weight - 500) / 100) * 500
	}

	fmt.Printf("Cost = %d\n", cost)
}
