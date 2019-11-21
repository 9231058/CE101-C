/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 21-11-2019
 * |
 * | File Name:     8-5.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	fmt.Println(check(10, 10, 90))
	fmt.Println(check(10, 20, 90))
	fmt.Println(check(10, 10, 120))
}

func check(side1, side2, angle float32) int {
	c := 0
	if side1 == side2 {
		c |= 1
	}
	if angle == 90 {
		c |= 2
	}
	return c
}
