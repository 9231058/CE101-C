/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 17-11-2017
 * |
 * | File Name:     p1.go
 * +===============================================
 */

package main

import "fmt"

func digit(n int) []int {
	if n == 0 {
		return make([]int, 0)
	}
	return append(digit(n/10), n%10)
}

func main() {
	var n int

	fmt.Scanf("%d", &n)

	ds := digit(n)
	for i := 0; i < len(ds)/2; i++ {
		if ds[i] != ds[len(ds)-1-i] {
			fmt.Println("False")
			return
		}
	}
	fmt.Println("True")
}
