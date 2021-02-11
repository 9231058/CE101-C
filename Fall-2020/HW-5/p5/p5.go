/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 04-01-2021
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

var reader = bufio.NewReader(os.Stdin)

func scanf(format string, variables ...interface{}) { fmt.Fscanf(reader, format, variables...) }

func main() {
	var n, k int
	scanf("%d %d", &n, &k)

	fmt.Println(josephus(n-1, k+1))

}

func josephus(n, k int) int {
	if n == 1 {
		return 1
	} else {
		return (josephus(n-1, k)+k-1)%n + 1
	}
}
