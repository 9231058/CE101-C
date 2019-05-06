/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 06-05-2019
 * |
 * | File Name:     p4.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	res := nextFibonacci(n)
	if res == -1 {
		fmt.Println("Error")
	} else {
		fmt.Printf("%d %d\n", res-n, res)
	}
}

func nextFibonacci(n int) int {
	i := 0
	res := 0
	for ; res < n; i++ {
		res = fibonacci(i)
	}

	if res == n {
		return fibonacci(i)
	}
	return -1
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
