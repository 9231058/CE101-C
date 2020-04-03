/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 22-03-2020
* |
* | File Name:     p3.go
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
	s := 0
	t := 0.0
	for i := 0; i < n; i++ {
		var student string
		scanf("%s\n", &student)
		switch student {
		case "H":
			s++
			t = math.Max(t, float64(s))
			break
		case "C":
			t = math.Max(t, float64(s))
			s = 0
			break
		}
	}
	fmt.Println(t)
}
