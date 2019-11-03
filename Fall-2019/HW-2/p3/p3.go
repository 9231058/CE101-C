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

	if _, err := fmt.Scanf("%f %f %f %f", &a, &b, &c, &d); err != nil {
		return
	}

	// caculates the (a) section of problem 3
	rA := a/(a+b) + b/(b+c) + c/(c+d)
	fmt.Printf("a: %.02f\n", rA)

	// caculates the (b) section of problem 3
	rB := math.Pow(math.E, math.Log2(math.Abs(a))+math.Sin(b)+math.Tan(c))
	fmt.Printf("b: %.02f\n", rB)

	// caculates the (c) section of problem 3
	rC := a * (b + a) * b * (math.Log(c))
	fmt.Printf("c: %.02f\n", rC)

	// caculates the (d) section of problem 3
	rD := math.Pow(1/math.Pow(a, 1.2)+2/math.Pow(b, 2.3), 4.5)
	fmt.Printf("d: %.02f\n", rD)
}
