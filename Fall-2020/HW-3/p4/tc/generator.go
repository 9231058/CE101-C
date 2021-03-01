package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i < 50; i++ {
		salary := 1000 * (rand.Float64() + 0.01)
		sal := int(salary)
		fmt.Printf("%d %.1f;", i, float64(sal))
	}
}
