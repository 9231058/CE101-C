/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 20-12-2017
 * |
 * | File Name:     p4.go
 * +===============================================
 */

package main

import (
	"container/list"
	"fmt"
)

func print(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Printf("\n")
}

func insertBefore(l *list.List, x, y interface{}, equal func(i, j interface{}) bool) {
	for e := l.Front(); e != nil; e = e.Next() {
		if equal(e.Value, y) {
			l.InsertBefore(x, e)
			return
		}
	}
	l.PushBack(x)
}

func main() {
	l := list.New()

	s := []struct{ x, y int }{
		{20, 100},
		{30, 100},
		{40, 30},
		{50, 20},
		{60, 10},
	}

	for _, s := range s {
		insertBefore(l, s.x, s.y, func(i, j interface{}) bool {
			if i.(int) == j.(int) {
				return true
			}
			return false
		})
		print(l)
	}
}
