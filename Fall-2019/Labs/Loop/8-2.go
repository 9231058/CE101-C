/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 12-11-2019
 * |
 * | File Name:     8-2.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	p := 1 << n

	var lsum float64
	var rsum float64

	for i := 1; i <= p; i++ {
		lsum += 1.0 / float64(i)
	}

	for i := p; i >= 1; i-- {
		rsum += 1.0 / float64(i)
	}

	fmt.Println(rsum == lsum)
}
