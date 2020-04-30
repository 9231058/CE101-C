/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 16-04-2020
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
	var a, b, c int
	scanf("%d %d %d\n", &a, &b, &c)
	if isPalindrome(convert(a, b, c)) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isPalindrome(s string) bool {
	l := len(s) - 1
	for i := range s {
		if s[i] != s[l-i] {
			return false
		}
	}
	return true
}
func convert(i int, base int, c int) string {
	num, _ := strconv.ParseInt(fmt.Sprintf("%d", i), base, 64)
	s := big.NewInt(num).Text(c)
	return s
}
