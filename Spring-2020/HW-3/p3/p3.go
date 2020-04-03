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
	var a, b, c float64
	scanf("%f %f %f\n", &a, &b, &c)
	ans := math.Inf(1)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				ta, tb, tc := (a + float64(i)), (b + float64(j)), (c + float64(k))
				sum := math.Abs(ta-tb) + math.Abs(ta-tc) + math.Abs(tc-tb)
				ans = math.Min(ans, sum)
			}
		}
	}
	fmt.Println(int(ans))
}
