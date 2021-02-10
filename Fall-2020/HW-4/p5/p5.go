/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 20-12-2020
 * |
 * | File Name:     p5.go
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
	var n, mod int64
	scanf("%d %d", &n, &mod)

	result := int64(1)
	for i := int64(0); i < n; i++ {
		result = (result * (i + 2)) % mod
	}

	if result == 0 {
		result = mod
	}

	fmt.Println(result)
}
