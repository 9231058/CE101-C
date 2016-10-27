/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 27-10-2016
 * |
 * | File Name:     p5.go
 * +===============================================
 */

package main

import (
	"fmt"
)

func main() {
	var y1, m1, d1 int
	var y2, m2, d2 int

	fmt.Scanf("%d %d %d", &y1, &m1, &d1)
	fmt.Scanf("%d %d %d", &y2, &m2, &d2)

	td1 := y1*360 + m1*30 + d1
	td2 := y2*360 + m2*30 + d2

	var td3 int

	if td1 > td2 {
		td3 = td1 - td2
	} else {
		td3 = td2 - td1
	}

	y3 := td3 / 360
	td3 %= 360
	m3 := td3 / 30
	td3 %= 30
	d3 := td3

	fmt.Printf("%d %d %d\n", y3, m3, d3)
}
