package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for j := 0; j < 100; j++ {
		n := rand.Intn(200) + 5
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", rand.Intn(250)+1)
		}
		fmt.Printf("0;")
	}
}
