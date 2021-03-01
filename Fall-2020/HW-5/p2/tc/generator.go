package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		n := rand.Intn(9) + 2

		fmt.Println(n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				v := rand.Intn(21) - 10
				fmt.Printf("%d ", v)
			}
			fmt.Printf("\n")
		}
		fmt.Printf(";")
	}
}
