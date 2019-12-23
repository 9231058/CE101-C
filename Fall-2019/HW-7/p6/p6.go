/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 23-12-2019
 * |
 * | File Name:     p6.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var n int

	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Scanf("%d", &nums[i]); err != nil {
			return
		}
	}

	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	numbers := make(map[int]bool)

	for _, num := range nums {
		numbers[num] = true
	}

	for i := 1; i <= len(nums); i++ {
		if numbers[i] == false {
			return i
		}
	}
	return len(nums) + 1
}
