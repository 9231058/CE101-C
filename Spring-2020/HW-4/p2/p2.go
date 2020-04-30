/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 16-04-2020
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

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var n int
	scanf("%d\n", &n)
	fmt.Println(climb(0, n))
}

func climb(i, n int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	return climb(i+1, n) + climb(i+2, n)
}

func climbStairs(n int) int {
	dp := make([]int, n+3)
	dp[0], dp[1], dp[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}
