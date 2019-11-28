package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 20; i++ {
		n := rand.Intn(10) + 10

		fmt.Printf("%d;", n)
	}
}
