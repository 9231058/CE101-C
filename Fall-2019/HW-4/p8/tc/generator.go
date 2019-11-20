package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 200; i++ {
		n := rand.Intn(1000*1000*1000) + 1000

		fmt.Printf("%d;", n)
	}
}
