/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 23-11-2018
 * |
 * | File Name:     p3.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	i := 1
	a := 1
	b := 2
	for i < n {
		if i%2 == 0 {
			b = a + b
		} else {
			a = a + b
		}
		i++
	}
	if n%2 == 0 {
		fmt.Println(b)
	} else {
		fmt.Println(a)
	}
}
