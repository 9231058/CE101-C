package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i <= 100; i++ {
		n := rand.Intn(500) + 1
		m := rand.Intn(500) + 1

		a := rand.Intn(50) + 1

		fmt.Printf("%d %d %d;", n, m, a)
	}
}
