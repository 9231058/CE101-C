/*
 * In The Name Of God
 * ========================================
 * [] File Name : p5.go
 *
 * [] Creation Date : 16-11-2015
 *
 * [] Created By : Elahe Jalalpour (el.jalalpour@gmail.com)
 * =======================================
 */
/*
 * Copyright (c) 2015 Elahe Jalalpour.
 */
package main

import (
	"fmt"
)

func main() {
	var direction int
	var n int
	var counter int = 0
	var run int = 0

	var new_number int = 0
	var old_number int = 0

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &direction)
	fmt.Scanf("%d", &new_number)

	for i := 1; i < n; i++ {
		old_number = new_number
		fmt.Scanf("%d", &new_number)
		if direction == 1 && new_number <= old_number {
			if counter > 0 {
				run++
			}
			counter = 0
		} else if direction == -1 && new_number >= old_number {
			if counter > 0 {
				run++
			}
			counter = 0
		} else {
			counter++
		}
	}
	if counter > 0 {
		run++
	}

	fmt.Println(run)
}
