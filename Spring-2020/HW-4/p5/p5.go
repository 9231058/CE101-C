/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 16-04-2020
* |
* | File Name:     p5.go
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
	var n, k int
	scanf("%d %d", &n, &k)
	fmt.Println(kthGrammar(n, k))
}

func kthGrammar(N int, K int) int {
	if N == 1 && K == 1 {
		return 0
	}
	k := math.Ceil(float64(K) / 2)
	kr := K % 2
	res := kthGrammar(N-1, int(k))
	if res == 1 {
		return kr
	} else {
		return 1 - kr
	}
}
