/*
 * +===============================================
 * | Author:        Elahe Jalalpour (el.jalalpour@gmail.com)
 * |
 * | Creation Date: 27-11-2015
 * |
 * | File Name:     p1.go
 * +===============================================
 */
package main

import "fmt"

func sequence_generator(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 3
	} else {
		return n + sequence_generator(n-1) + sequence_generator(n-2)
	}

}
func check(x int) int {
	for i := 1; sequence_generator(i) <= x; i++ {
		if sequence_generator(i) == x {
			return i
		}
	}
	return -1
}
func print_between(a int, b int) {
	for i := a; i < b; i++ {
		if check(i) != -1 {
			fmt.Printf("%d @ %d\n", i, check(i))
		}
	}
}

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	print_between(a, b)
}
