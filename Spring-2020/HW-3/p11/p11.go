/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 22-03-2020
* |
* | File Name:     p11.go
* +===============================================
 */

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var n int
	scanf("%d\n", &n)
	h, v := 0, 0
	for i := 0; i < n; i++ {
		var c string
		scanf("%s\n", &c)
		switch c {
		case "U":
			h++
			break
		case "D":
			h--
			break
		case "L":
			v--
			break
		case "R":
			v++
			break
		}
	}
	if v == 0 && h == 0 {
		fmt.Println("YES")
	} else {
		ho := int(math.Abs(float64(h)))
		ve := int(math.Abs(float64(v)))
		fmt.Printf("y axis and x axis distances are : %d %d\n", ho, ve)
	}
}
