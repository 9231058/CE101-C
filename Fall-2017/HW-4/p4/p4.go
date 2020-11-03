/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 11-11-2017
 * |
 * | File Name:     p4.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)

	if a%(b+1) == 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
