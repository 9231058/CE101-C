/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 26-11-2016
 * |
 * | File Name:     p6.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d\n", &n)

	fmt.Printf("%d\n", tile(n))
}

func tile(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return tile(n-1) + tile(n-2)
	}
}
