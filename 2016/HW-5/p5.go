/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 27-11-2015
 * |
 * | File Name:     p5.go
 * +===============================================
 */
package main

import "fmt"

func main() {
	var a1, b1, c1, a2, b2, c2 int
	var op byte
	fmt.Scanf("%d %d %d\n", &a1, &b1, &c1)
	fmt.Scanf("%d %d %d\n", &a2, &b2, &c2)
	fmt.Scanf("%c\n", &op)

	switch op {
	case '+':
		fmt.Printf("0 * x^4 + 0 * x^3 + %d * x^2 + %d * x^1 + %d\n", a1+a2, b1+b2, c1+c2)
	case '-':
		fmt.Printf("0 * x^4 + 0 * x^3 + %d * x^2 + %d * x^1 + %d\n", a1-a2, b1-b2, c1-c2)
	case '*':
		fmt.Printf("%d * x^4 + %d * x^3 + %d * x^2 + %d * x^1 + %d\n", a1*a2, a2*b1+a1*b2, a2*c1+a1*c2+b1*b2, b2*c1+b1*c2, c1*c2)
	}

}
