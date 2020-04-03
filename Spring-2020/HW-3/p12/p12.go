/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 22-03-2020
* |
* | File Name:     p12.go
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
	start := int(math.Pow(10, float64(n)-1))
	end := int(math.Pow(10, float64(n)))
	flag := false
	for i := start; i < end; i++ {
		tempI := i
		flag = true
		for tempI > 0 {
			digit := int(tempI) % 10
			if digit == 0 || int(i)%digit == 0 {
				flag = false
				break
			}
			tempI /= 10
		}
		if flag {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("no corona")
}
