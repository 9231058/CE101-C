package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		t := rand.Intn(5)
		if t < 4 {
			a1 := rand.Intn(1000) - 500
			a2 := rand.Intn(1000) - 500
			a3 := rand.Intn(1000) - 500
			fmt.Printf("%d %d %d;", a1, a2, a3)
		} else {
			a12 := rand.Intn(1000) - 500
			a3 := rand.Intn(1000) - 500
			fmt.Printf("%d %d %d;", a12, a12, a3)
		}
	}
}
