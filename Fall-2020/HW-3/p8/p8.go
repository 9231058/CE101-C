/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 05-12-2020
 * |
 * | File Name:     p8.go
 * +===============================================
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var start, end string
	if _, err := fmt.Scanf("%s\n%s\n", &end, &start); err != nil {
		return
	}

	sStart := strings.Split(start, ":")
	sEnd := strings.Split(end, ":")

	sH, _ := strconv.Atoi(sStart[0])
	eH, _ := strconv.Atoi(sEnd[0])
	sM, _ := strconv.Atoi(sStart[1])
	eM, _ := strconv.Atoi(sEnd[1])
	sS, _ := strconv.Atoi(sStart[2])
	eS, _ := strconv.Atoi(sEnd[2])

	seconds := 0
	if eH > sH {
		seconds = (eH-sH)*3600 + (eM-sM)*60 + eS - sS
	} else {
		help := 24 * 3600
		seconds = (eH-sH)*3600 + (eM-sM)*60 + eS - sS + help
	}

	hours := seconds / 3600
	seconds = seconds % 3600

	minutes := seconds / 60
	seconds = seconds % 60

	fmt.Printf("%.2d:%.2d:%.2d", hours, minutes, seconds)
}
