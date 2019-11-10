/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 10-11-2019
 * |
 * | File Name:     p7.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	fmt.Println(poorPigs(n, 15, 15))
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	count := minutesToTest/minutesToDie + 1

	pigs := math.Log(float64(buckets)) / math.Log(float64(count))

	return int(math.Ceil(pigs))
}
