/*
*In The Name of God
*
* +===============================================
* | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
* |
* | Creation Date: 14-05-2020
* |
* | File Name:     p2.go
* +===============================================
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")
	subString, _ := reader.ReadString('\n')
	subString = strings.TrimRight(subString, "\n")
	replacement, _ := reader.ReadString('\n')
	replacement = strings.TrimRight(replacement, "\n")

	newLine := strings.ReplaceAll(line, subString, replacement)

	fmt.Println(newLine)
}
