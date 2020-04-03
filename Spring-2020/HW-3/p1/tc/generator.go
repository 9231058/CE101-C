package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		a := rand.Intn(2000000) - 1000000
		b := rand.Intn(2000000) - 1000000
		k := rand.Intn(500000)

		fmt.Printf("%d %d %d;", a, b, k)

	}
}
