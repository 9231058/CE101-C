package main

import "fmt"

func main() {
	var n, m int

	fmt.Scanf("%d %d", &n, &m)

	for a := n; a <= m; a++ {
		for b := n; b <= m; b++ {
			for c := n; c <= m; c++ {
				if a*a == b*b+c*c {
					fmt.Printf("%d %d %d\n", a, b, c)
				}
			}
		}
	}
}
