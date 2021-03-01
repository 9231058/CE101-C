/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 05-12-2020
 * |
 * | File Name:     p5.go
 * +===============================================
 */

package main

import (
	"fmt"
)

func main() {
	var x1, y1, x2, y2, x3, y3 int
	if _, err := fmt.Scanf("%d %d\n%d %d\n%d %d", &x1, &y1, &x2, &y2, &x3, &y3); err != nil {
		return
	}

	mAB := float64(y2-y1) / float64(x2-x1)
	mBC := float64(y3-y2) / float64(x3-x2)

	if mAB*mBC == -1 || ((x3-x2 == 0 && y1-y2 == 0) || (x1-x2 == 0 && y3-y2 == 0)) {
		fmt.Printf("%d %d\n", x1-x2+x3, y1-y2+y3)
	} else {
		fmt.Printf("Not Possible")
	}

}
