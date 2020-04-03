/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 22-03-2020
* |
* | File Name:     p6.go
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
	for i := 2; i < n; i++ {

		if i == sumOfDivisors(sumOfDivisors(i)) && i != sumOfDivisors(i) && i < sumOfDivisors(i) {
			fmt.Printf("%d-%d\n", i, sumOfDivisors(i))
		}
	}
}

func sumOfDivisors(n int) int {
	res := 0
	for i := 1; i <= int(math.Pow(float64(n), 0.5)); i++ {
		if n%i == 0 {
			if i == n/i || i == 1 {
				res += i
			} else {
				res += i + n/i
			}
		}
	}
	return res
}
