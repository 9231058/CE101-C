/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 22-03-2020
* |
* | File Name:     p10.go
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
	var a, b int
	scanf("%d %d\n", &a, &b)
	if b >= a || (a-b)%2 != 0 {
		fmt.Println("Not Possible")
		return
	}
	for i := 1; i <= a; i++ {
		for j := 1; j <= a; j++ {
			if i > (a-b)/2 && i <= a-((a-b)/2) && j > (a-b)/2 && j <= a-((a-b)/2) {
				fmt.Print("  ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}
