package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func isPrime(n int) bool {
	bound := int(math.Sqrt(float64(n)))

	for i := 2; i < bound+1; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
