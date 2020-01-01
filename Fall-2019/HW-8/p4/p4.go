/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 30-12-2019
 * |
 * | File Name:     p4.go
 * +===============================================
 */

package main

import (
	"fmt"
	"sort"
)

type student struct {
	gpa  int
	name string
	id   int
}

func main() {
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}

	students := make([]student, n)

	for i := 0; i < n; i++ {
		var fname, lname string
		if _, err := fmt.Scanf("%s %s", &fname, &lname); err != nil {
			return
		}
		students[i].name = fname + " " + lname

		if n, err := fmt.Scanf("%d", &students[i].id); err != nil {
			fmt.Println(n, err)
			return
		}
		if _, err := fmt.Scanf("%d", &students[i].gpa); err != nil {
			return
		}
	}

	sort.Slice(students, func(i, j int) bool {
		if students[i].gpa == students[j].gpa {
			return students[i].id < students[j].id
		}
		return students[i].gpa > students[j].gpa
	})

	for i := 0; i < n; i++ {
		fmt.Printf("%s\n", students[i].name)
		fmt.Printf("%d\n", students[i].id)
		fmt.Printf("%d\n", students[i].gpa)
	}
}
