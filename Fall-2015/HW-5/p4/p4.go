/*
 * +===============================================
 * | Author:        Elahe Jalalpour (el.jalalpour@gmail.com)
 * |
 * | Creation Date: 27-11-2015
 * |
 * | File Name:     p4.go
 * +===============================================
 */
package main

import "fmt"

func main() {
	Coordinate()
}

func Coordinate() {
	var n int
	fmt.Scanf("%d", &n)
	switch n % 4 {
	case 0:
		fmt.Printf("(%d,%d)\n", -(n / 4), n/4)
	case 1:
		fmt.Printf("(%d,%d)\n", -(n / 4), -(n / 4))
	case 2:
		fmt.Printf("(%d,%d)\n", n/4+1, -(n / 4))
	case 3:
		fmt.Printf("(%d,%d)\n", n/4+1, n/4+1)
	}
}
