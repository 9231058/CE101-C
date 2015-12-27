package main

import "fmt"
import "sort"

type Date struct {
	Year  int
	Month int
	Day   int
}
type DateArray []Date

func (array DateArray) Len() int {
	return len(array)
}

func (array DateArray) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

func (array DateArray) Less(i, j int) bool {
	if array[i].Year == array[j].Year {
		if array[i].Month == array[j].Month {
			return array[i].Day < array[j].Day
		} else {
			return array[i].Month < array[j].Month
		}
	} else {
		return array[i].Year < array[j].Year
	}
}

func main() {
	var num int
	var dir int

	fmt.Scanf("%d", &num)

	var dates []Date

	for i := 0; i <= num; i++ {
		dates = append(dates, Date{})
		fmt.Scanf("%d %d %d", &dates[i].Year, &dates[i].Month, &dates[i].Day)
	}
	fmt.Scanf("%d", &dir)

	sort.Sort(DateArray(dates))

	switch dir {
	case 1:
		for i := 0; i < num; i++ {
			fmt.Printf("%d/%d/%d\n", dates[i].Year, dates[i].Month, dates[i].Day)
		}
	case -1:
		for i := num - 1; i >= 0; i-- {
			fmt.Printf("%d/%d/%d\n", dates[i].Year, dates[i].Month, dates[i].Day)
		}
	}
}
