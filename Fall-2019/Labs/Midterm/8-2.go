/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 21-11-2019
 * |
 * | File Name:     8-2.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var x float64 = 1          // students must set the initial value and test to see its effect
	for i := 0; i < 100; i++ { // students must set the loop counter and test to see its effect
		x = x - f(x)/fp(x)
	}
	fmt.Println(x)
}

func f(x float64) float64 {
	return x*x + 3*x + 2
}

func fp(x float64) float64 {
	return 2*x + 3
}
