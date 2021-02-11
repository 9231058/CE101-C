package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		n := rand.Intn(500) + 10
		k := rand.Intn(n-5) + 1
		fmt.Printf("%d %d;", n, k)
	}
}
