package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		l1, l2 := rand.Intn(30)+3, rand.Intn(30)+3
		fmt.Printf("%d ", rand.Intn(9)+1)
		for j := 0; j < l1; j++ {
			fmt.Printf("%d ", rand.Intn(10))
		}
		fmt.Printf("-1\n")

		fmt.Printf("%d ", rand.Intn(9)+1)
		for j := 0; j < l2; j++ {
			fmt.Printf("%d ", rand.Intn(10))
		}
		fmt.Printf("-1\n;")
	}

}
