/*
 * In The Name Of God
 * ========================================
 * [] File Name : p3.go
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

func number_3_digits_test(n int) bool {
	return (n > 99 && n < 1000)
}

func power_3_digits_test(n int) bool {
	n1 := n % 10
	n2 := (n / 10) % 10
	n3 := (n / 100) % 10

	return (n1*n1*n1+n2*n2*n2+n3*n3*n3 == n)
}

func main() {
	var n int
	var m int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)
	if number_3_digits_test(n) && number_3_digits_test(m) {
		for i := n; i <= m; i++ {
			if power_3_digits_test(i) {
				fmt.Printf("%d\n", i)
			}
		}
	} else {
		fmt.Printf("Error\n")
	}
}
