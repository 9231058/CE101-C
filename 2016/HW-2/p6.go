/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 27-10-2016
 * |
 * | File Name:     p6.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var n, m int
	fmt.Scanf("%d %d", &m, &n)

	for i := 0; i < 4; i++ {
		fmt.Println(r.Intn(n-m) + m)
	}
}
