/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 11-11-2019
 * |
 * | File Name:     8-3.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
