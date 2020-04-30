/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 16-04-2020
* |
* | File Name:     p3.go
* +===============================================
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func determinan(q1, q4, q3, q2 float64) float64 {
	return (q1 * q4) - (q2 * q3)
}

func mult(b, c []float64) []float64 {
	d := make([]float64, 9)
	for i := 0; i < 3; i++ {
		var s float64
		s = 0
		for j := 0; j < 3; j++ {
			s = b[i*3]*c[j] + b[i*3+1]*c[j+3] + b[i*3+2]*c[j+6]
			d[i*3+j] = s
		}
	}
	return d
}
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	a := make([]float64, 9)
	b := make([]float64, 9)
	c := make([]float64, 9)
	for j := 0; j < 9; j++ {
		scanf("%f ", &a[j])
	}
	scanf("\n")
	for j := 0; j < 9; j++ {
		scanf("%f ", &b[j])
	}
	scanf("\n")
	for j := 0; j < 9; j++ {
		scanf("%f ", &c[j])
	}
	scanf("\n")
	det := a[0]*determinan(a[4], a[8], a[7], a[5]) - a[1]*determinan(a[3], a[8], a[6], a[5]) + a[2]*determinan(a[3], a[7], a[6], a[4])
	detainvers := 1 / det

	for i := 0; i < 9; i++ {
		b[i] *= detainvers
	}

	d := mult(b, c)
	for i := 0; i < 9; i++ {
		fmt.Printf("%.2f ", d[i])
	}
}
