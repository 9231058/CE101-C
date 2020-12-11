package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	a = -5.0
	b = math.Pi / 6
	c = math.Pi / 5

	result := math.Pow(math.E, math.Log2(math.Abs(a))+math.Sin(b)+math.Tan(c))

	fmt.Printf("%.2f\n", result)

	a = 3.5
	b = 2
	c = 9

	result = a * (b + a) * b * math.Log(c)

	fmt.Printf("%.2f\n", result)

	a = 5
	b = 6

	result = math.Pow((1/math.Pow(a, 0.3) + (2 / math.Pow(b, 2.3))), 4.5)

	fmt.Printf("%.2f\n", result)
}
