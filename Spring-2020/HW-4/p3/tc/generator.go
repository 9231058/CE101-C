package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 20; i++ {
		s := ""
		for j := 0; j < 3; j++ {
			for l := 0; l < 9; l++ {
				s += fmt.Sprintf("%d ", rand.Intn(20))
			}
			s += "\n"
		}
		fmt.Printf("%s;", s)
	}
}
