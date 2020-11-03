/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 03-11-2017
 * |
 * | File Name:     p4.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	if n < 10000 || n > 99999 {
		fmt.Printf("Error\n")
		return
	}

	a := n % 10
	n /= 10
	b := n % 10
	n /= 10
	n /= 10
	d := n % 10
	n /= 10
	e := n % 10

	if a == e && b == d {
		fmt.Printf("True\n")
	} else {
		fmt.Printf("False\n")
	}

}
