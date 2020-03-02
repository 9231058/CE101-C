/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 29-02-2020
 * |
 * | File Name:     p2.go
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
	var a, b, r, q int
	var c, m float32
	scanf("%d%d%f%d%f%d", &a, &b, &c, &r, &m, &q)
	z := ((8.8 * float32(a+b) * 2 / c) - 0.5 + float32(2*a/(q+r))) / (float32(a+b) * float32(1/m))
	x := (float32(-b) + float32(b*b) + 2.4*float32(a)*c) / float32(2*a)

	fmt.Printf("Result for \"Z\" is %.2f\nResult for \"X\" is %.3f\n", z, x)
}
