/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-10-2019
 * |
 * | File Name:     p2.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}

	// stores the reverse of n in m and number of digits in nd
	nd := 0
	m := 0
	for n != 0 {
		m = m*10 + n%10
		n = n / 10
		nd++
	}

	r := 0
	for m != 0 {
		nd--
		d := m % 10
		if d == 0 {
			r = r * 10
		}
		r = r*10 + d
		m = m / 10
	}
	for nd > 0 {
		r = r * 100
		nd--
	}

	fmt.Println(r)
}
