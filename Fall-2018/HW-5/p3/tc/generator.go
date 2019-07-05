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
		n := rand.Intn(20) + 1

		fmt.Printf("%d;", n)
	}
}
