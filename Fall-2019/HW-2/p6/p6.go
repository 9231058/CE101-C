package main

import "fmt"

func main() {
	var w, h, k int

	if _, err := fmt.Scanf("%d %d %d", &w, &h, &k); err != nil {
		return
	}

	r := 0
	for i := 0; i < k; i++ {
		c := 2*(w+h) - 4
		r += c
		h -= 4
		w -= 4
	}

	fmt.Println(r)
}
