package main

import "fmt"

func main() {
	var input string

	fmt.Scanf("%s", &input)

	counts := make(map[rune]int)

	var sorted []rune
	var output []rune

	for _, c := range input {
		if _, ok := counts[c]; !ok {
			sorted = append(sorted, c)
		}
		counts[c]++
	}

	for _, c := range sorted {
		if counts[c] > 1 {
			output = append(output, c)
		}
	}

	fmt.Printf("%s\n", string(output))
}
