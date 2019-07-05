/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 21-12-2018
 * |
 * | File Name:     p2.go
 * +===============================================
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	intersection := make(map[int]bool)

	var n1, n2 int
	fmt.Scanf("%d", &n1)
	for i := 0; i < n1; i++ {
		var v int
		fmt.Scanf("%d", &v)
		intersection[v] = true
	}
	fmt.Scanf("%d", &n2)
	for i := 0; i < n2; i++ {
		var v int
		fmt.Scanf("%d", &v)
		intersection[v] = true
	}

	// sorts the keys of intersection set
	var keys []int
	for k := range intersection {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// lets show the results
	fmt.Println(len(keys))
	for _, k := range keys {
		fmt.Print(k, " ")
	}
	fmt.Println()

}
