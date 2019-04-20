package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	var index int
	var even, odd int
	for a != 0 {
		if index%2 == 0 {
			even += a % b
		} else {
			odd += a % b
		}
		index++
		a /= b
	}

	if even == odd {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
