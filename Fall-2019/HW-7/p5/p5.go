package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}
	set := make([]int, n)

	for i := 0; i < n; i++ {
		if _, err := fmt.Scanf("%d", &set[i]); err != nil {
			return
		}
	}

	subsets := make([][]int, 0)
	for i := 0; i < 1<<n; i++ {
		subset := make([]int, 0)
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				subset = append(subset, set[j])
			}
		}
		subsets = append(subsets, subset)
	}

	sort.Slice(subsets, func(i, j int) bool {
		if len(subsets[i]) != len(subsets[j]) {
			return len(subsets[i]) < len(subsets[j])
		}
		for k := 0; k < len(subsets[i]); k++ {
			if subsets[i][k] != subsets[j][k] {
				return subsets[i][k] < subsets[j][k]
			}
		}
		return false
	})

	for _, s := range subsets {
		fmt.Printf("{ ")
		for _, e := range s {
			fmt.Printf("%d ", e)
		}
		fmt.Printf("}\n")
	}
}
