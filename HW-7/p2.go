package main

import "fmt"

func switch_func(a *int, b [11]int, c int) *int {
	a = &b[c]
	return a
}

func main() {
	var c int
	fmt.Println("Enter c")
	fmt.Scanf("%d", &c)
	var arr [11]int
	for i := 1; i < 11; i++ {
		arr[i] = i
	}
	var parr *int
	parr = switch_func(parr, arr, c)
	fmt.Println(*parr)
}
