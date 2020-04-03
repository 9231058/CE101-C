/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 23-03-2020
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
	var ans string
	start := 0
	end := 99999
	mid := (start + end) / 2
	fmt.Println(mid)
	for scanf("%s\n", &ans); ans != "EQ"; scanf("%s\n", &ans) {
		switch ans {
		case "H":
			start = mid
			break
		case "L":
			end = mid
			break
		}
		mid = (start + end) / 2
		fmt.Println(mid)
	}
	fmt.Println("Fatabarak allah ahsan alkhaleghin")
}
