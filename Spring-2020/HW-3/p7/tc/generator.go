package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		b := rand.Intn(7) + 1
		c := rand.Intn((b*9)-1) + 1
		fmt.Printf("%d %d;", b, c)
	}
}
