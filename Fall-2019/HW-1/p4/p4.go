/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-10-2019
 * |
 * | File Name:     p4.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}

	var p uint64 = 1
	var i uint64
	for i = 1; i <= uint64(n); i++ {
		p = (p * i) / gcd(p, i)
	}

	fmt.Println(p)
}

func gcd(a, b uint64) uint64 {
	for b > 0 {
		r := a % b
		a = b
		b = r
	}
	return a
}
