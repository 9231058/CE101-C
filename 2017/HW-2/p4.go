/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-10-2017
 * |
 * | File Name:     p4.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var x int

	fmt.Scanf("%d", &x)

	fmt.Printf("Perimeter: %d\n", 4*x)
	fmt.Printf("Area: %d\n", x*x)
	fmt.Printf("Diameter: %.2f\n", math.Sqrt(2)*float64(x))
	fmt.Printf("Inside Circle: %.2f\n", 3.14*float64(x*x)*0.25)
	fmt.Printf("Outside Circle: %.2f\n", 3.14*math.Sqrt(2)*float64(x))
}
