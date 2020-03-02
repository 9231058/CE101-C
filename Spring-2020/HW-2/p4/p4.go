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
	var x1, x2, a, b int
	scanf("%d%d%d%d", &x1, &x2, &a, &b)
	// if we assume moving to right as positive movement
	// we have for 1: x = at + x1
	// then for 2: x = bt + x2
	// so we must solve : at + x1 = bt + x2 -> t(a-b) = (x2-x1)
	t := (x2 - x1) / (a - b)
	x := a*t + x1
	fmt.Printf("They met in %d after %d jumps and lived happily ever after :)\n", x, t)
}
