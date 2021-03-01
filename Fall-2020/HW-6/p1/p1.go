/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 17-01-2021
 * |
 * | File Name:     p1.go
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

type Node struct {
	data int
	next *Node
	prev *Node
}

func main() {
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	t := strings.Split(line, " ")
	var cur *Node
	cur = nil
	for i := range t {
		number, _ := strconv.Atoi(t[i])
		if number == 0 {
			break
		}
		n := &Node{
			data: number,
			next: nil,
			prev: nil,
		}
		if cur == nil {
			cur = n
		} else {
			n.prev = cur
			cur.next = n
			cur = n
		}
	}

	for n := cur; n != nil; n = n.prev {
		fmt.Printf("%d ", n.data)
	}

}
