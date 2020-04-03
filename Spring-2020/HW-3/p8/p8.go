/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 22-03-2020
* |
* | File Name:     p9.go
* +===============================================
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var i, j, n int

	// Input number of rows from user
	scanf("%d", &n)

	// Print the upper part of the arrow
	for i = 1; i < n; i++ {
		// Print trailing (2*rownumber-2) spaces
		for j = 1; j <= (2*i - 2); j++ {
			fmt.Printf(" ")
		}

		// Print inverted right triangle star pattern
		for j = i; j <= n; j++ {
			fmt.Printf("*")
		}

		fmt.Println("")
	}

	// Print lower part of the arrow
	for i = 1; i <= n; i++ {
		// Print trailing (2*n - 2*rownumber) spaces
		for j = 1; j <= (2*n - 2*i); j++ {
			fmt.Printf(" ")
		}

		// Print simple right triangle star pattern
		for j = 1; j <= i; j++ {
			fmt.Printf("*")
		}

		fmt.Println("")
	}

}
