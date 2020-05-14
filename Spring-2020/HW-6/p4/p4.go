/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 14-05-2020
* |
* | File Name:     p4.go
* +===============================================
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var n int64
	scanf("%d\n", &n)
	x := make([]int, n)
	for i := int64(0); i < n; i++ {
		scanf("%d ", &x[i])
	}
	sort.Ints(x)
	sb := new(strings.Builder)
	for i := range x {
		sb.WriteString(fmt.Sprintf("%d,", x[i]))
	}
	fmt.Println(sb.String())
}
