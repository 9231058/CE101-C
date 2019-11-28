package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		n := rand.Intn(1000) + 10

		fmt.Printf("%d;", n)
	}
}
