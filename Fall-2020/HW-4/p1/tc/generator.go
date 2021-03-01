package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		n := rand.Intn(10000)

		a := new(strings.Builder)
		b := new(strings.Builder)
		for j := 0; j < n; j++ {
			bit := rand.Intn(2)
			a.WriteString(fmt.Sprintf("%d", bit))
			bit = rand.Intn(2)
			b.WriteString(fmt.Sprintf("%d", bit))
		}

		fmt.Printf("%d\n%s\n%s;", n, a.String(), b.String())
	}
}
