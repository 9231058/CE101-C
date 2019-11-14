/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 14-11-2019
 * |
 * | File Name:     p5.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var q, d int
	var n int64

	fmt.Scanf("%d %d\n", &q, &n)
	for i := 0; i < q; i++ {
		var s byte
		var c int64
		fmt.Scanf("%c %d\n", &s, &c)
		if s == '-' {
			if n >= c {
				n -= c
			} else {
				d++
			}
		} else {
			n += c
		}
	}

	fmt.Printf("%d %d\n", n, d)
}
