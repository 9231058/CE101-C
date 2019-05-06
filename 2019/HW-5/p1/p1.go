package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	var index = 1
	res := digit(n, index)
	for res != -1 {
		fmt.Printf("%d ", res)
		index++
		res = digit(n, index)
	}
	fmt.Print("\n")
	for index > 1 {
		index--
		fmt.Printf("%d ", digit(n, index))
	}
	fmt.Print("\n")
}

func digit(n, m int) int {
	i := 0

	for n != 0 {
		i++
		if i == m {
			return n % 10
		}
		n /= 10
	}

	return -1
}
