/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 09-11-2018
 * |
 * | File Name:     7-3.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n, k int64

	fmt.Scanf("%d %d", &n, &k)

	d := n / k
	if d%2 == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
