/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 30-10-2017
 * |
 * | File Name:     p5.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	fmt.Printf("%d = 1 / 8 + 6 / 7\n", 1/8+6/7)
	fmt.Printf("%d = 1 + 8 / 6 / 7\n", 1+8/6/7)
	fmt.Printf("%d = 1 + 8 / 6 %% 7\n", 1+8/6%7)
	fmt.Printf("%d = 1 + 8 - 6 %% 7\n", 1+8-6%7)
	fmt.Printf("%d = 1 + 8 * 6 / 7\n", 1+8*6/7)
	fmt.Printf("%d = 1 * 8 + 6 / 7\n", 1*8+6/7)
}
