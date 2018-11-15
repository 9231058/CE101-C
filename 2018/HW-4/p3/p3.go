/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 15-11-2018
 * |
 * | File Name:     p3.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	for n != 0 {
		d := n % 10
		fmt.Printf("%d: ", d)
		for i := 0; i < d; i++ {
			fmt.Print(d)
		}
		fmt.Println()
		n = n / 10
	}
}
