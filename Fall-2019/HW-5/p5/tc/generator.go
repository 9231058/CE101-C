package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 150; i++ {
		n := rand.Intn(1000) + 25

		fmt.Printf("%d;", n)
	}
}
