package main

import (
	"fmt"
)

func main() {
	fmt.Println("sdfsdf")
	data := []int{1, 3, 5}
	data1 := []int{2, 4, 6}
	fmt.Println(merge(data, data1))
	data = []int{10, 2, 4, 3, 90, 23, 2, 7}
	fmt.Println(merge_sort(data))
}

func merge_sort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	middle := len(data) / 2

	left := merge_sort(data[middle:])
	right := merge_sort(data[:middle])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	res := make([]int, len(left)+len(right))
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res[i+j] = left[i]
			i++
		} else {
			res[i+j] = right[j]
			j++
		}
	}

	for i < len(left) {
		res[i+j] = left[i]
		i++
	}

	for j < len(right) {
		res[i+j] = right[j]
		j++
	}

	return res
}
