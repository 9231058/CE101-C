package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("%d %d;", 3, 500)

	for i := 0; i < 20; i++ {
		n := rand.Intn(1000*1000*1000) + 1000
		k := rand.Intn(1000 * 1000 * 1000)

		fmt.Printf("%d %d;", n, k)
	}
}
