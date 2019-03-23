package main

import "fmt"

func main() {
	var x, s1, m1, s2, m2 int

	fmt.Scanf("%d %d %d %d %d", &x, &s1, &m1, &s2, &m2)

	fmt.Println(x - (s2*60+m2-(s1*60+m1))%x)
}
