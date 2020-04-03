package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 20; i++ {
		a := rand.Intn(18) + 2
		b := rand.Intn(a - 2)

		fmt.Printf("%d %d;", a, b)

	}
}
