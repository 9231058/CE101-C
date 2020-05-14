package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		sb := new(strings.Builder)
		n := rand.Intn(100000)
		sb.WriteString(fmt.Sprintf("%d\n", n))
		for j := 0; j < n; j++ {
			x := rand.Intn(100000) - 50000
			sb.WriteString(fmt.Sprintf("%d ", x))
		}
		sb.WriteString(";")
		fmt.Print(sb.String())
	}
}
