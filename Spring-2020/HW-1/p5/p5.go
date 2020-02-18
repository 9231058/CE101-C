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
	// 9^5 = 59049
	// 6 * (9^5) = 354294 < 10^7
	// so there is no answer greater than 354294
	var res int64
	for i := 2; i < int(math.Pow(10, 6)); i++ {
		if i == getFifthSumOfDigits(i) {
			res += int64(i)
		}
	}
	fmt.Println(res)
}
func getFifthSumOfDigits(n int) int {
	sum := 0
	for i := n; i > 0; i /= 10 {
		sum += int(math.Pow(float64(i%10), 5))
	}
	return sum
}
