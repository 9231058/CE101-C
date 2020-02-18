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
	length := int(math.Log10(float64(n))) + 1
	power := length - 1
	tempN := n
	for i := power; i >= 0; i-- {
		digit := n / int(math.Pow(10, float64(i)))
		if digit == 6 {
			tempN += 3 * (int(math.Pow(10, float64(i))))
			break
		}
		n = n % int(math.Pow(10, float64(i)))
	}
	fmt.Printf("%d\n", tempN)
}
