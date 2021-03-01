package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for k := 0; k < 20; k++ {
		l := rand.Intn(50) + 5
		arr := make([]int, l)
		for i := range arr {
			arr[i] = rand.Intn(500)
		}

		seqLen := rand.Intn(l-3) + 3
		index := rand.Intn(l - seqLen)

		l2 := rand.Intn(20) + 2
		replacement := make([]int, l2)
		for i := range replacement {
			replacement[i] = rand.Intn(100)
		}

		for i := range arr {
			fmt.Printf("%d ", arr[i])
		}
		fmt.Printf("-1\n")

		for i := index; i < index+seqLen; i++ {
			fmt.Printf("%d ", arr[i])
		}
		fmt.Printf("-1\n")

		for i := range replacement {
			fmt.Printf("%d ", replacement[i])
		}
		fmt.Printf("-1\n;")
	}
}
