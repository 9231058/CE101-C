/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 29-02-2020
 * |
 * | File Name:     p7.go
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
	scanf("%d", &n)
	res := make([]int, 6)
	carry := 0
	for i := 0; n != 0; i++ {
		digit := n % 10
		digit = digit + carry + 1
		carry = digit / 10
		res[i] = digit % 10
		n /= 10
	}
	res[5] = carry
	for i := range res {
		fmt.Print(res[i])
	}
}
