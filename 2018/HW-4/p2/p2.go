/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 15-11-2018
 * |
 * | File Name:     p2.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n, m, a int

	fmt.Scanf("%d %d %d", &n, &m, &a)

	// found idicates the leftmost place that digit `a` happens in `n`
	found := 0
	// digit counter
	d := 1

	for n != 0 {
		if n%10 == a {
			found = d
		}
		d++
		n = n / 10
	}

	// there is a digit `a` in `n` and it happens in the `m` digits from left
	if found != 0 && found > d-m {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
