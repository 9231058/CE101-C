/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 05-11-2017
 * |
 * | File Name:     p3.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var x00, x01, x10, x11 int
	var y00, y01, y10, y11 int

	fmt.Scanf("%d %d", &x01, &y01)

	fmt.Scanf("%d %d", &x11, &y11)

	fmt.Scanf("%d %d", &x10, &y10)

	fmt.Scanf("%d %d", &x00, &y00)

	var s, r, d, p bool

	if (y11-y01)*(x10-x00) == (y10-y00)*(x11-x01) {
		if (y01-y00)*(x11-x10) == (y11-y10)*(x01-x00) {
			p = true
			if (x01-x00)*(x11-x01)+(y11-y01)*(y01-y00) == 0 {
				r = true
			}
			if (y01-y00)*(y01-y00)+(x01-x00)*(x01-x00) == (y11-y01)*(y11-y01)+(x11-x01)*(x11-x01) {
				d = true
			}
			if r && d {
				s = true
			}
		}
	}

	fmt.Printf("Square: %v\n", s)
	fmt.Printf("Rectangle: %v\n", r)
	fmt.Printf("Parallelogram: %v\n", p)
	fmt.Printf("Diamond: %v\n", d)
}
