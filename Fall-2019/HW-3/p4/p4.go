/*
*
* In The Name of God
*
* +===============================================
* | Author:        Parham Alvani <parham.alvani@gmail.com>
* |
* | Creation Date: 10-11-2019
* |
* | File Name:     p4.go
* +===============================================
 */

package main

import "fmt"

func main() {
	var x1, y1 int
	if _, err := fmt.Scanf("%d %d", &x1, &y1); err != nil {
		return
	}

	var x2, y2 int
	if _, err := fmt.Scanf("%d %d", &x2, &y2); err != nil {
		return
	}

	var x3, y3 int
	if _, err := fmt.Scanf("%d %d", &x3, &y3); err != nil {
		return
	}

	var x4, y4 int
	if _, err := fmt.Scanf("%d %d", &x4, &y4); err != nil {
		return
	}

	p := 0
	if length(x1, y1, x2, y2) == length(x2, y2, x3, y3) {
		p++
	}
	if length(x1, y1, x3, y3) == length(x2, y2, x4, y4) {
		p++
	}
	if length(x1, y1, x2, y2) == length(x3, y3, x4, y4) {
		p++
	}
	if length(x2, y2, x3, y3) == length(x4, y4, x1, y1) {
		p++
	}

	if p == 4 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func length(x1, y1, x2, y2 int) int {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}
