/*
 * +===============================================
 * | Author:        Elahe Jalalpour (el.jalalpour@gmail.com)
 * |
 * | Creation Date: 27-11-2015
 * |
 * | File Name:     p1.go
 * +===============================================
 */
package main

import "fmt"

func Rec_func(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + Rec_func(n-1)
	}

}
func Check(x int) int {
	for i := 1; Rec_func(i) <= x; i++ {
		if Rec_func(i) == x {
			return i
		}
	}
	return -1
}
func Show_bet(a int, b int) {
	for i := a; i < b; i++ {
		if Check(i) != -1 {
			fmt.Printf("%d @ %d\n", i, Check(i))
		}
	}
}

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	Show_bet(a, b)
}
