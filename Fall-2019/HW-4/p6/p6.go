/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 14-11-2019
 * |
 * | File Name:     p6.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int
	var g int
	var w int

	fmt.Scanf("%d %d\n", &n, &g)
	for i := 0; i < n; i++ {
		var h int
		fmt.Scanf("%d", &h)
		if h > g {
			w += 2
		} else {
			w++
		}
	}

	fmt.Println(w)
}
