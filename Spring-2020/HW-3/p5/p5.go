/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 22-03-2020
* |
* | File Name:     p5.go
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
	var year int
	fmt.Scanln(&year)
	if year%400 == 0 {
		fmt.Println(year, "True")
	} else if year%100 == 0 {

		fmt.Println(year, "False")
	} else if year%4 == 0 {
		fmt.Println(year, "True")
	} else {
		fmt.Println(year, "False")
	}
}
