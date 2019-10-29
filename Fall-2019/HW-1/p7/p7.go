/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-10-2019
 * |
 * | File Name:     p7.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}

	m := make(map[int]bool)
	for i := 2; i*i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if j%i == 0 {
				m[j] = true
			}
		}
	}

	for i := 2; i <= n; i++ {
		if m[i] == false {
			fmt.Println(i)
		}
	}
}
