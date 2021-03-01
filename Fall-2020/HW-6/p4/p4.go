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
	"math"
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
	var firstHead *Node
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
			firstHead = n
		} else {
			n.prev = firstList
			firstList.next = n
			firstList = n
		}
	}

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	t = strings.Split(line, " ")
	var secondList, secondHead *Node
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
			secondHead = n
		} else {
			n.prev = secondList
			secondList.next = n
			secondList = n
		}
	}

	var resultList, resultHead *Node
	for {
		firstDigit, secondDigit := -1, -1

		if secondHead == nil && firstHead == nil {
			break
		}

		if firstHead != nil {
			firstDigit = firstHead.data
		}

		if secondHead != nil {
			secondDigit = secondHead.data
		}

		if firstDigit > secondDigit {
			firstHead = firstHead.next
		} else if secondDigit > firstDigit {
			secondHead = secondHead.next
		} else if secondDigit != -1 {
			firstHead = firstHead.next
		}

		n := &Node{
			data: int(math.Max(float64(firstDigit), float64(secondDigit))),
			next: nil,
			prev: nil,
		}

		if resultList == nil {
			resultList = n
			resultHead = n
		} else {
			n.prev = resultList
			resultList.next = n
			resultList = n
		}
	}

	for n := resultHead; n != nil; n = n.next {
		fmt.Printf("%d ", n.data)
	}

}
