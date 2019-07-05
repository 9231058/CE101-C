package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 20; i++ {
		x := rand.Intn(10) + 1
		s1 := rand.Intn(4) + 1
		m1 := rand.Intn(60)
		s2 := rand.Intn(5) + s1
		m2 := rand.Intn(60)

		fmt.Printf("%d %d %d %d %d;", x, s1, m1, s2, m2)
	}
}
