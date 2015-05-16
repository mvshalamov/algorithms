package main

import (
	"fmt"
)

func main() {
	data := []int{5, 7, 3, 19, 4}
	//fmt.Println(partiton(data, 0, 0, len(data)-1))
	quicksort(data, 0, len(data)-1)
	fmt.Println(data)
}

func quicksort(data []int, first_index int, last_index int) {
	if first_index < last_index {

		pivot := partiton(data, first_index, first_index, last_index)
		quicksort(data, first_index, pivot-1)
		quicksort(data, pivot+1, last_index)
	}
}

func partiton(data []int, pivot_index int, first_index int, last_index int) int {
	left := first_index
	right := last_index
	pivot_value := data[pivot_index]
	go_while := true

	for go_while {
		for left <= right && pivot_value > data[left] {
			left = left + 1
		}

		for left <= right && pivot_value < data[right] {
			right = right - 1
		}

		if right <= left {
			break
		}

		tmp := data[left]
		data[left] = data[right]
		data[right] = tmp

	}

	return right

}
