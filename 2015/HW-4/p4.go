/*
 * In The Name Of God
 * ========================================
 * [] File Name : p4.go
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
	"math"
)

func n_th_power_algorithm(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n%2 == 1 {
		return x * n_th_power_algorithm(x, n)
	} else {
		tmp := n_th_power_algorithm(x, n/2)
		return tmp * tmp
	}
}

func n_th_root_algorithm(A float64, x float64, n int) float64 {
	nf := float64(n)
	return (1 / nf) * ((nf-1)*x + (A / n_th_power_algorithm(x, n-1)))
}

func main() {
	var A float64
	var n int

	fmt.Scanf("%f", &A)
	fmt.Scanf("%d", &n)

	old_x := 0.0
	new_x := A
	e := 0.0001

	for math.Abs(new_x-old_x) > e {
		old_x = new_x
		new_x = n_th_root_algorithm(A, old_x, n)
	}
	fmt.Printf("%.3f\n", new_x)
}
