package main

import (
	"fmt"
	"math/rand"
)

func main() {
	pi := 3.141592

	for i := 0; i < 50; i++ {
		max_deg := rand.Intn(179) + 1
		deg := rand.Intn(max_deg) + 1

		degree := float64(deg) * pi / 180
		max_degree := float64(max_deg) * pi / 180

		invalid := rand.Int() % 10
		l1, l2 := rand.Intn(100), rand.Intn(100)
		if invalid > 8 {
			fmt.Printf("%d %d %f %f;", l1, l2, max_degree, degree)
		} else {
			fmt.Printf("%d %d %f %f;", l1, l2, degree, max_degree)
		}
	}
}
