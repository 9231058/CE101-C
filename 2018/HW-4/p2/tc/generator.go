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
		n := rand.Intn(1000) + 1000 // [0 1000) + 1000 = [1000 2000)
		m := rand.Intn(5)           // [0 5)
		a := rand.Intn(10)          // [0 10)

		fmt.Printf("%d %d %d;", n, m, a)
	}
}
