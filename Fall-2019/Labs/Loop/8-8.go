/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 12-11-2019
 * |
 * | File Name:     8-8.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	fmt.Println(findNthDigit(n))
}

func findNthDigit(n int) int {
	i := 0
	var s int64
	var pre int
	for s < int64(n) {
		pre = int(s)
		i++
		s += nDigitNumbers(i)
	}
	n -= (pre + 1)

	t := pow10(i-1) + n/i

	d := i - (n % i)

	p := 1
	for t > 0 {
		if p == d {
			return t % 10
		}
		t = t / 10
		p++
	}

	return 0
}

func pow10(n int) int {
	r := 1

	for i := 0; i < n; i++ {
		r *= 10
	}

	return r
}

func nDigitNumbers(n int) int64 {
	var r int64 = 9

	for i := 1; i < n; i++ {
		r *= 10
	}

	return r * int64(n)
}
