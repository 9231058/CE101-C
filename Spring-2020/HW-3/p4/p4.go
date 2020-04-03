/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 22-03-2020
* |
* | File Name:     p4.go
* +===============================================
 */
package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var a int
	scanf("%d\n", &a)
	a, _ = strconv.Atoi(big.NewInt(int64(a)).Text(7))
	n := a
	var rev, dig int
	f := 1
	for n > 0 {
		n /= 10
		f *= 10
	}
	n = a
	f /= 10
	for i := 0; n > 0; i++ {
		dig = (n % 10) * f
		rev += dig
		n /= 10
		f /= 10
	}

	if rev == a {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
