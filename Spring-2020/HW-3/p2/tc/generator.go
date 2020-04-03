package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		a1 := rand.Intn(100000000)
		a2 := rand.Intn(100000000)
		a3 := rand.Intn(100000000)
		fmt.Printf("%d %d %d;", a1, a2, a3)
	}
}
