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
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var a, b, k float64
	scanf("%f %f %f\n", &a, &b, &k)
	aLow := math.Floor(a/k) * k
	aHigh := math.Floor(a/k+1) * k
	bLow := math.Floor(b/k) * k
	bHigh := math.Floor(b/k+1) * k
	p1 := math.Abs(a-aLow) + math.Abs(aLow-bLow)/k + math.Abs(b-bLow)
	p2 := math.Abs(a-aLow) + math.Abs(aLow-bHigh)/k + math.Abs(b-bHigh)
	p3 := math.Abs(a-aHigh) + math.Abs(aHigh-bLow)/k + math.Abs(b-bLow)
	p4 := math.Abs(a-aHigh) + math.Abs(aHigh-bHigh)/k + math.Abs(b-bHigh)
	min := math.Min(math.Abs(a-b), math.Min(math.Min(p1, p2), math.Min(p3, p4)))
	fmt.Println(int(min))
}
