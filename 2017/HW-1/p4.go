/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-10-2017
 * |
 * | File Name:     2017/HW-1/p4.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var x, y int

	fmt.Scanf("%d", &x)

	t := x

	for t != 0 {
		y = y*10 + (t % 10)
		t = t / 10
	}

	fmt.Println(x - y)
}
