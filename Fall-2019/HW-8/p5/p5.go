/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 31-12-2019
 * |
 * | File Name:     p5.go
 * +===============================================
 */

package main

import "fmt"

type phone struct {
	no    string
	phone string
}

var phones = make([]phone, 0)

func main() {
	for {
		choice := menu()
		switch choice {
		case 1:
			var ph phone
			fmt.Printf("No: ")
			if _, err := fmt.Scanf("%s", &ph.no); err != nil {
				return
			}
			fmt.Printf("Phone: ")
			if _, err := fmt.Scanf("%s", &ph.phone); err != nil {
				return
			}
			phones = append(phones, ph)
		case 2:
			var no string
			fmt.Printf("No: ")
			if _, err := fmt.Scanf("%s", &no); err != nil {
				return
			}

			index := -1
			for i := range phones {
				if phones[i].no == no {
					index = i
					break
				}
			}
			if index == -1 {
				fmt.Printf("%s does not exist\n", no)
			} else {
				phones = append(phones[:index], phones[index+1:]...)
				fmt.Printf("%s has been removed\n", no)
			}
		case 3:
			for i, ph := range phones {
				fmt.Printf("%d\n", i+1)
				fmt.Printf("No: %s\n", ph.no)
				fmt.Printf("Phone: %s\n", ph.phone)
			}
		case 4:
			return
		}
	}
}

func menu() int {
	fmt.Println("1) Add")
	fmt.Println("2) Remove")
	fmt.Println("3) List")
	fmt.Println("4) Quit")

	var choice int
	if _, err := fmt.Scanf("%d", &choice); err != nil {
		return 4
	}

	return choice
}
