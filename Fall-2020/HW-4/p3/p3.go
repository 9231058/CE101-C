/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 20-12-2020
 * |
 * | File Name:     p3.go
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

func scanf(format string, variables ...interface{}) { fmt.Fscanf(reader, format, variables...) }

func main() {
	var n int

	scanf("%d\n", &n)

	lines := make([]string, n)

	for i := range lines {
		lines[i], _ = reader.ReadString('\n')
		lines[i] = strings.TrimSpace(lines[i])
	}

	for i := range lines {
		s := new(strings.Builder)

		for j := range lines[i] {
			if j == 0 {
				s.WriteString(strings.ToUpper(string(lines[i][j])))
				continue
			}

			if lines[i][j-1] == byte(' ') {
				s.WriteString(strings.ToUpper(string(lines[i][j])))
				continue
			}

			s.WriteString(strings.ToLower(string(lines[i][j])))
		}

		lines[i] = s.String()
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
