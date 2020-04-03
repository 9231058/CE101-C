/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 22-03-2020
* |
* | File Name:     p7.go
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
	var b, c int
	scanf("%d %d\n", &b, &c)
	start := int(math.Pow(10, float64(b-1)))
	end := int(math.Pow(10, float64(b))) - 1
	max := 0
	for i := end; i >= start; i-- {
		if sumOfDigits(i) == c {
			max = i
			fmt.Println(max)
			return
		}
	}
}

func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit
		n /= 10
	}
	return sum
}
