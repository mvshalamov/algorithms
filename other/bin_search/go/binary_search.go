package main

import (
	"fmt"
)

func main() {
	data := []int{1, 3, 6, 8, 9, 11, 14, 17}
	fmt.Println(binarysearch(data, 9))
}

func binarysearch(data []int, item int) int {
	first := 0
	last := len(data) - 1
	mid := -1
	if len(data) != 0 {
		for first <= last {
			mid := (first + last) / 2
			if data[mid] == item {
				return mid
			} else if item < data[mid] {
				last = mid - 1
			} else {
				first = mid + 1
			}
		}
		return mid
	}
	return -1
}
