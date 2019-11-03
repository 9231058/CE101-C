package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		w := rand.Intn(100) + 3
		h := rand.Intn(100) + 3

		min := h
		if min > w {
			min = w
		}
		k := rand.Intn(min/4 + 1)

		fmt.Printf("%d %d %d;", w, h, k)
	}
}
