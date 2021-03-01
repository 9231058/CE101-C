package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	for i := 0; i < 100; i++ {
		l1 := rand.Intn(60) + 3
		arr1 := make([]int, l1)
		for i := range arr1 {
			arr1[i] = rand.Intn(100)
		}

		l2 := rand.Intn(60) + 3
		arr2 := make([]int, l2)
		for i := range arr2 {
			arr2[i] = rand.Intn(100)
		}

		sort.Ints(arr1)
		sort.Ints(arr2)

		for i := range arr1 {
			fmt.Printf("%d ", arr1[len(arr1)-i-1])
		}
		fmt.Printf("-1\n")
		for i := range arr2 {
			fmt.Printf("%d ", arr2[len(arr2)-i-1])
		}
		fmt.Printf("-1\n;")
	}
}
