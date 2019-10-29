/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-10-2019
 * |
 * | File Name:     p1.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var a, b int

	if _, err := fmt.Scanf("%d %d", &a, &b); err != nil {
		return
	}

	for b > 0 {
		r := a % b
		a = b
		b = r
	}
	fmt.Println(a)
}
