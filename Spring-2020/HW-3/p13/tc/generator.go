package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		n := rand.Intn(991) + 10
		fmt.Println(n + 1)
		fmt.Println("C")
		for j := 0; j < n; j++ {
			k := rand.Intn(3) + 1
			if k%3 == 0 {
				fmt.Println("C")
			} else {
				fmt.Println("H")
			}
		}
		fmt.Printf(";")
	}
}
