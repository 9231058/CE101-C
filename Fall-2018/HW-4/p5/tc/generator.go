package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("%d %d;", 1, 10)
	fmt.Printf("%d %d;", 5, 10)

	for i := 0; i < 20; i++ {
		m := rand.Intn(100)
		n := rand.Intn(m)

		fmt.Printf("%d %d;", n, m)
	}
}
