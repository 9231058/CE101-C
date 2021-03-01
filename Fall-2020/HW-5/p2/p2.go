/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 04-01-2021
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

var reader = bufio.NewReader(os.Stdin)

func scanf(format string, variables ...interface{}) { fmt.Fscanf(reader, format, variables...) }

func main() {
	var n int
	scanf("%d\n", &n)
	matrix := make([][]int64, n)

	for i := range matrix {
		matrix[i] = make([]int64, n)
		for j := range matrix[i] {
			scanf("%d", &matrix[i][j])
		}
		scanf("\n")
	}

	fmt.Println(determinan(n, matrix))
}

func calc_cofactor(n int, dst, src [][]int64, index int) {
	x := 0
	y := 0
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if j != index {
				dst[y][x] = src[i][j]
				x++
				y += x / (n - 1)
				x %= n - 1
			}
		}
	}
}

func determinan(n int, matrix [][]int64) int64 {
	if n == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}
	cof := int64(1)
	s := int64(0)
	for i := 0; i < n; i++ {
		cofactor := make([][]int64, n-1)
		for i := range cofactor {
			cofactor[i] = make([]int64, n-1)
		}

		calc_cofactor(n, cofactor, matrix, i)
		s += cof * matrix[0][i] * determinan(n-1, cofactor)
		cof *= -1
	}
	return s
}
