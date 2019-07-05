/*
 * In The Name Of God
 * ========================================
 * [] File Name : p1.go
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

func main() {
	var n int
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i += 2 {
		sec := (2*n - 2*i) / 4
		for j := 0; j < sec; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		for j := 0; j < 2*sec; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		//for j := 0; j < sec; j++ {
		//	fmt.Print(" ")
		//}
		fmt.Println()
	}
	for i := n - 2; i >= 1; i -= 2 {
		sec := (2*n - 2*i) / 4
		for j := 0; j < sec; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		for j := 0; j < 2*sec; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		//for j := 0; j < sec; j++ {
		//	fmt.Print(" ")
		//}
		fmt.Println()
	}

}
