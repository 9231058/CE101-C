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
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var n int
	scanf("%d\n", &n)
	for i := 1; i < n; i++ {
		if check(i) {
			fmt.Printf("%d ", i)
		}
	}
}

func check(n int) bool {
	sum := 0
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
		if sum > n {
			return true
		}
	}
	return false
}
