/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 17-01-2021
 * |
 * | File Name:     p4.go
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

type Node struct {
	data int
	next *Node
	prev *Node
}

func main() {
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	line = strings.Replace(line, "-1", "", 1)

	find, _ := reader.ReadString('\n')
	find = strings.TrimSpace(find)
	find = strings.Replace(find, "-1", "", 1)

	replacement, _ := reader.ReadString('\n')
	replacement = strings.TrimSpace(replacement)
	replacement = strings.Replace(replacement, "-1", "", 1)

	// fuck this shit, I'm out
	line = strings.Replace(line, find, replacement, -1)

	fmt.Println(line)
}
