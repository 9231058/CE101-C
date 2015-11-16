/*
 * In The Name Of God
 * ========================================
 * [] File Name : p2.go
 *
 * [] Creation Date : 16-11-2015
 *
 * [] Created By : Elahe Jalalpour (el.jalalpour@gmail.com)
 * =======================================
 */
/*
 * Copyright (c) 2015 Elahe Jalalpour.
 */

package main

import (
	"fmt"
)

func reverse(n int) (int, int) {
	result := 0
	digit_count := 0

	for ; n != 0; n = n / 10 {
		result = result*10 + n%10
		digit_count++
	}
	return result, digit_count
}

func digits(n int) {
	d := 0
	n, d = reverse(n)

	for ; d != 0; n = n / 10 {
		fmt.Printf("%d: ", n%10)
		for i := 0; i < n%10; i++ {
			fmt.Printf("%d", n%10)
		}
		fmt.Println()
		d--
	}
}

func main() {
	var input int

	fmt.Scanf("%d", &input)
	digits(input)
}
