/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 15-11-2018
 * |
 * | File Name:     p4.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int
	var digits = make([]int, 0)

	fmt.Scanf("%d", &n)

	for n != 0 {
		d := n % 10
		digits = append(digits, d)
		n = n / 10
	}

	for left, right := 0, len(digits)-1; left < right; left, right = left+1, right-1 {
		digits[left], digits[right] = digits[right], digits[left]
	}

	for _, d := range digits {
		fmt.Printf("%d: ", d)
		for i := 0; i < d; i++ {
			fmt.Print(d)
		}
		fmt.Println()
	}

}
