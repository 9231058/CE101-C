package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("1 2 3 4 0;")
	fmt.Printf("1 2 3 4 5 0;")

	for i := 0; i < 20; i++ {
		n := rand.Intn(100)
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", rand.Intn(1024))
		}
		fmt.Printf("0;")
	}
}
