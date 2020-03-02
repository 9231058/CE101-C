/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 29-02-2020
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
	var bmi float64
	var weight int
	var height float64
	scanf("%d %f\n", &weight, &height)
	bmi = float64(weight) / (height * height)
	fmt.Printf("%.2f\n", bmi)
}
