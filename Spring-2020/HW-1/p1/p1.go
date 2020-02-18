/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 18-02-2020
 * |
 * | File Name:     p1.go
 * +===============================================
 */

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var n int
	scanf("%d\n", &n)
	for i := 1; i < n; i++ {
		if isSD(i, getDigits(i)) {
			fmt.Printf("%d ", i)
		}
	}
}

// check if i is Self Dividing
func isSD(i int, digits []int) bool {
	for _, v := range digits {
		if v == 0 {
			return false
		}
		if (i % v) != 0 {
			return false
		}
	}
	return true
}

func getDigits(n int) []int {
	// get length of n
	digits := make([]int, int(math.Log10(float64(n)))+1)
	x := 0
	for i := n; i > 0; i /= 10 {
		digits[x] = i % 10
		x++
	}
	return digits
}
