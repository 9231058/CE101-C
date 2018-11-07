/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 08-11-2018
 * |
 * | File Name:     p5.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var a1, b1 int
	var a2, b2 int

	fmt.Scanf("%d %d", &a1, &b1)
	fmt.Scanf("%d %d", &a2, &b2)

	// based on catalan numbers, these segments have 2 forms and following intersections:
	// ( ( ) ) x4
	// a1 a2 b1 b2
	// a1 a2 b2 b1
	// a2 a1 b2 b1
	// a2 a1 b1 b2
	// ( ) ( ) x2
	// a2 b2 a1 b1
	// a1 b1 a2 b2

	if a1 < a2 {
		if b1 < a2 { // a1 b1 a2 b2
			fmt.Println(0)
		} else {
			if b1 < b2 { // a1 a2 b1 b2
				fmt.Println(b1 - a2)
			} else { // a1 a2 b2 b1
				fmt.Println(b2 - a2)
			}
		}
	} else {
		if b2 < a1 { // a2 b2 a1 b1
			fmt.Println(0)
		} else {
			if b1 < b2 { // a2 a1 b1 b2
				fmt.Println(b1 - a1)
			} else { // a2 a1 b2 b1
				fmt.Println(b2 - a1)
			}
		}
	}
}
