/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-10-2017
 * |
 * | File Name:     p2.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	fmt.Println(sum(10, 11))
	fmt.Println(sub(20, 18))
}

func addOne(x int) int {
	return x + 1
}

func subOne(x int) int {
	return x - 1
}

func sum(x, y int) int {
	if y == 0 {
		return x
	}

	return addOne(sum(x, subOne(y)))
}

func sub(x, y int) int {
	if y == 0 {
		return x
	}

	return subOne(sub(x, subOne(y)))
}
