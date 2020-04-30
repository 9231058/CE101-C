package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	for i := 1; i <= 30; i++ {
		k := rand.Intn(int(math.Pow(2, float64(i-1)))) + 1
		fmt.Printf("%d %d;", i, k)
	}
}
