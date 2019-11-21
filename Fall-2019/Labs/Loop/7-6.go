/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 11-11-2019
 * |
 * | File Name:     7-6.go
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
	if sp == math.Round(sp) {
		for i := 2; i < int(sp); i++ {
			if int(sp)%i == 0 {
				fmt.Println("false")
				return
			}
		}
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
