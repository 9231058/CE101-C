/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 15-11-2018
 * |
 * | File Name:     generator.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 20; i++ {
		a1 := rand.Intn(10)
		b1 := rand.Intn(10)
		c1 := rand.Intn(10)

		a2 := rand.Intn(10)
		b2 := rand.Intn(10)
		c2 := rand.Intn(10)

		fmt.Printf("%d %d %d\n%d %d %d;", a1, b1, c1, a2, b2, c2)
	}
}
