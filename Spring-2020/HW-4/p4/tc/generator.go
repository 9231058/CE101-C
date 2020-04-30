package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

func main() {
	for i := 0; i < 250; i++ {
		a := rand.Intn(1000) + 1
		b := (rand.Intn(9) + 2)
		c := (rand.Intn(9) + 2)
		s := big.NewInt(int64(a)).Text(b)
		fmt.Printf("%s %d %d;", s, b, c)
	}
}
