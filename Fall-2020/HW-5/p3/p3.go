/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 04-01-2021
 * |
 * | File Name:     p3.go
 * +===============================================
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readline() string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func main() {
	key := readline()
	n, _ := strconv.Atoi(readline())
	result := 0

	for i := 0; i < n; i++ {
		line := readline()
		if strings.Contains(line, key) {
			result++
			continue
		}

		found := true

		for c := range key {
			index := strings.Index(line, string(key[c]))

			if index != -1 {
				line = line[index+1:]
				continue
			}

			found = false
			break
		}

		if found {
			result++
		}
	}

	fmt.Println(result)
}
