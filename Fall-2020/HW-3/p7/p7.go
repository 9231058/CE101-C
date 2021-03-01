/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 05-12-2020
 * |
 * | File Name:     p7.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m, a int
	if _, err := fmt.Scanf("%d %d %d", &n, &m, &a); err != nil {
		return
	}

	x := math.Ceil(float64(n) / float64(a))
	y := math.Ceil(float64(m) / float64(a))

	result := int(x * y)

	fmt.Println(result)
}
