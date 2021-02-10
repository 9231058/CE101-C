/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 20-12-2020
 * |
 * | File Name:     p1.go
 * +===============================================
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func scanf(format string, variables ...interface{}) { fmt.Fscanf(reader, format, variables...) }

func main() {
	var n int
	var a, b string

	scanf("%d\n%s\n%s", &n, &a, &b)

	result := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}

		if i != len(a)-1 && (a[i] == b[i+1] && a[i+1] == b[i]) {
			i++
			result++
			continue
		}

		result++
	}

	fmt.Println(result)
}
