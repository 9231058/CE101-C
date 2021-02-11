/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 17-01-2021
 * |
 * | File Name:     p2.go
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
	var firstList *Node
	firstList = nil
	for i := range t {
		number, _ := strconv.Atoi(t[i])
		if number == -1 {
			break
		}
		n := &Node{
			data: number,
			next: nil,
			prev: nil,
		}
		if firstList == nil {
			firstList = n
		} else {
			n.prev = firstList
			firstList.next = n
			firstList = n
		}
	}

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	t = strings.Split(line, " ")
	var secondList *Node
	secondList = nil
	for i := range t {
		number, _ := strconv.Atoi(t[i])
		if number == -1 {
			break
		}
		n := &Node{
			data: number,
			next: nil,
			prev: nil,
		}
		if secondList == nil {
			secondList = n
		} else {
			n.prev = secondList
			secondList.next = n
			secondList = n
		}
	}

	var resultList *Node
	carry := 0
	for {
		var firstDigit, secondDigit int

		if secondList == nil && firstList == nil {
			if carry != 0 {
				n := &Node{
					data: (firstDigit + secondDigit + carry) % 10,
					next: nil,
					prev: nil,
				}

				n.prev = resultList
				resultList.next = n
				resultList = n
			}

			break
		}

		if firstList != nil {
			firstDigit = firstList.data
			firstList = firstList.prev
		}

		if secondList != nil {
			secondDigit = secondList.data
			secondList = secondList.prev
		}

		n := &Node{
			data: (firstDigit + secondDigit + carry) % 10,
			next: nil,
			prev: nil,
		}

		carry = (firstDigit + secondDigit + carry) / 10

		if resultList == nil {
			resultList = n
		} else {
			n.prev = resultList
			resultList.next = n
			resultList = n
		}
	}

	for n := resultList; n != nil; n = n.prev {
		fmt.Printf("%d ", n.data)
	}

}
