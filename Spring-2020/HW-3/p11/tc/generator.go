package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		num := rand.Intn(91) + 10
		fmt.Println(num)
		for j := 0; j < num; j++ {
			moves := []string{"U", "D", "L", "R"}
			n := rand.Intn(4)
			fmt.Println(moves[n])
		}
		fmt.Print(";")
	}
}
