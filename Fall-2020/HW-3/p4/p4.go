/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 05-12-2020
 * |
 * | File Name:     p4.go
 * +===============================================
 */

package main

import (
	"fmt"
)

func main() {
	var b float64

	var a int
	if _, err := fmt.Scanf("%d %f", &a, &b); err != nil {
		return
	}

	monthly_payment := 0.94 * b

	reward := float64(a/12) * (2 * b)

	insurance := float64(a/6) * 0.2 * b

	income := (monthly_payment * float64(a)) + reward - insurance

	fmt.Printf("%.2f\n", income)
}
