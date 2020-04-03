/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 21-03-2020
* |
* | File Name:     p2.go
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
	var a, b, c, ans float64
	scanf("%f %f %f \n", &a, &b, &c)
	if a == b && a == c && c == b {
		ans = 0

	} else {
		av := (a + b + c)
		av = av / 3
		if a == av || b == av || c == av {
			ans = 1
		} else {
			ans = 2
		}

	}
	fmt.Println(int(ans))
}
