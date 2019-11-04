/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 05-11-2019
 * |
 * | File Name:     6-5.go
 * +===============================================
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println(intToRoman(n))
}

func intToRoman(num int) string {
	var r strings.Builder
	thousands := num / 1000
	num = num % 1000
	switch thousands {
	case 1:
		r.WriteString("M")
	case 2:
		r.WriteString("MM")
	case 3:
		r.WriteString("MMM")
	}

	hundreds := num / 100
	num = num % 100
	switch hundreds {
	case 1:
		r.WriteString("C")
	case 2:
		r.WriteString("CC")
	case 3:
		r.WriteString("CCC")
	case 4:
		r.WriteString("CD")
	case 5:
		r.WriteString("D")
	case 6:
		r.WriteString("DC")
	case 7:
		r.WriteString("DCC")
	case 8:
		r.WriteString("DCCC")
	case 9:
		r.WriteString("CM")
	}

	decimal := num / 10
	num = num % 10
	switch decimal {
	case 1:
		r.WriteString("X")
	case 2:
		r.WriteString("XX")
	case 3:
		r.WriteString("XXX")
	case 4:
		r.WriteString("XL")
	case 5:
		r.WriteString("L")
	case 6:
		r.WriteString("LX")
	case 7:
		r.WriteString("LXX")
	case 8:
		r.WriteString("LXXX")
	case 9:
		r.WriteString("XC")
	}

	one := num
	switch one {
	case 1:
		r.WriteString("I")
	case 2:
		r.WriteString("II")
	case 3:
		r.WriteString("III")
	case 4:
		r.WriteString("IV")
	case 5:
		r.WriteString("V")
	case 6:
		r.WriteString("VI")
	case 7:
		r.WriteString("VII")
	case 8:
		r.WriteString("VIII")
	case 9:
		r.WriteString("IX")
	}

	return r.String()
}
