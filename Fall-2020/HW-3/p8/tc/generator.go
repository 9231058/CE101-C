package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 50; i++ {
		hour1 := rand.Intn(24)
		hour2 := rand.Intn(24)

		min1 := rand.Intn(60)
		min2 := rand.Intn(60)

		sec1 := rand.Intn(60)
		sec2 := rand.Intn(60)

		fmt.Printf("%.2d:%.2d:%.2d\n%.2d:%.2d:%.2d;", hour1, min1, sec1, hour2, min2, sec2)
	}
}
