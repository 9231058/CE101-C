/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 21-12-2018
 * |
 * | File Name:     p3.go
 * +===============================================
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	words := make(map[string]int)

	// number lines and each line maximum length
	var nLine, maxLen int
	fmt.Scanf("%d %d", &nLine, &maxLen)

	for i := 0; i < nLine; i++ {
		// reads stdin lines
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		// parses each line and stores its words
		for _, word := range strings.Fields(line) {
			words[word]++
		}
	}

	// sorts the keys of words dictionary
	var keys []string
	for k := range words {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// lets show the results
	fmt.Println(len(keys))
	for _, k := range keys {
		fmt.Println(k, words[k])
	}
}
