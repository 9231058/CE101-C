/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 21-04-2019
 * |
 * | File Name:     p1.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n, h int

	fmt.Scanf("%d %d", &h, &n)

	for i := 0; i < h; i++ {
		spaces := ((2*h - 1) - (2*i + 1)) / 2
		stars := 2*i + 1

		for j := 0; j < n; j++ {
			for k := 0; k < spaces; k++ {
				fmt.Print(" ")
			}
			for k := 0; k < stars; k++ {
				fmt.Print("*")
			}

			for k := 0; k < spaces; k++ {
				fmt.Print(" ")
			}

			fmt.Print(" ")
		}

		fmt.Println()
	}
}
