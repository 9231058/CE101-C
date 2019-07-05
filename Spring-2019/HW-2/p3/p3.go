/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 25-10-2018
 * |
 * | File Name:     p3.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64

	fmt.Scanf("%f %f %f %f", &a, &b, &c, &d)

	// caculates the (a) section of problem 3
	resultA := (a+b)/(2*a+3*b) + (b+c)/(2*b+3*c) + (c+d)/(2*c+3*d)
	fmt.Printf("a: %.02f\n", resultA)

	// caculates the (b) section of problem 3
	resultB := math.Pow(10, math.Log10(math.Abs(a+math.Sin(b)+math.Tan(c))))
	fmt.Printf("b: %.02f\n", resultB)

	// caculates the (c) section of problem 3
	resultC := a / math.Log(b+math.Sin(a)) * b * (math.Log(c))
	fmt.Printf("c: %.02f\n", resultC)

	// caculates the (d) section of problem 3
	resultD := math.Pow(100/math.Pow(a, 10.2)+200/math.Pow(b, 20.3), 40.5)
	fmt.Printf("d: %.02f\n", resultD)
}
