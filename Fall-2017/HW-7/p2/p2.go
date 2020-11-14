/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 20-12-2017
 * |
 * | File Name:     p2.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var m int
	var n int

	fmt.Scanf("%d", &m)

	A := make([][]int, 0)

	for {
		input := make([]int, m)
		var isZero = true

		for i := 0; i < m; i++ {
			fmt.Scanf("%d", &input[i])
			if input[i] != 0 {
				isZero = false
			}
		}
		if isZero {
			break
		}
		A = append(A, input)
		n++
	}
	fmt.Println("A:")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", A[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	B := make([][]int, m)
	for i := 0; i < m; i++ {
		B[i] = make([]int, n)
		for j := 0; j < n; j++ {
			B[i][j] = A[j][i]
		}
	}
	fmt.Println("B:")
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", B[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
		for j := 0; j < n; j++ {
			for k := 0; k < m; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	fmt.Println("C:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", C[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	D := make([][]int, m)
	for i := 0; i < m; i++ {
		D[i] = make([]int, m)
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				D[i][j] += B[i][k] * A[k][j]
			}
		}
	}
	fmt.Println("D:")
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", D[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

}
