/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 29-02-2020
 * |
 * | File Name:     p2.go
 * +===============================================
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var n int
	scanf("%X", &n)
	var b strings.Builder
	for n != 0 {
		// in each step. remainder is a number between 0 to 25 (26 numbers)
		// so we can cover all A-Z alphabets with %26 by adding the remainder value to 64
		// a problem is that Z has the value of 26, so 26%26=0
		// yet another problem is having 0 as the remainder. 0+64 = 64 (@ in asci). we need to cover all A-Z
		// if we decrement value of n. we can solve this issue as for Z's remainder (26%26=0) now we 25
		// also if we change the starting position of asci codes to 'A'(65), we can cover all A-Z (65)-(90)
		n--
		b.WriteString(string(byte(n%26 + 65)))
		n /= 26
	}
	res := b.String()
	b.Reset()
	for i := range res {
		b.WriteByte(res[len(res)-i-1])
	}
	fmt.Println(b.String())
}
