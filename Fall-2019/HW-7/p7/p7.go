/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 23-12-2019
 * |
 * | File Name:     p7.go
 * +===============================================
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	sub, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	sub = strings.TrimSuffix(sub, "\n")

	fmt.Printf("%s", strings.Replace(text, sub, "", 1))
}
