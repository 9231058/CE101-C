package main

import "fmt"

func main() {
	var x, s1, m1, s2, m2 int

	fmt.Scanf("%d %d %d %d %d", &x, &s1, &m1, &s2, &m2)

	diff := (s2*60 + m2 - (s1*60 + m1))

	if diff%x == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(x - diff%x)
	}
}
