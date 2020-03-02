/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 29-02-2020
 * |
 * | File Name:     p2.go
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
	var lemon, apple, pear int
	scanf("%d%d%d", &lemon, &apple, &pear)
	lemons := lemon / 1
	apples := apple / 2
	pears := pear / 4

	min := int(math.Min(math.Min(float64(lemons), float64(apples)), float64(pears)))
	extraLemons := lemon - min
	extraApples := apple - (min * 2)
	extraPears := pear - (min * 4)
	fmt.Printf("COMPOTE: %d\nEXTRA-LEMON: %d\nEXTRA-APPLE: %d\nEXTRA-PEAR: %d", min, extraLemons, extraApples, extraPears)
}
