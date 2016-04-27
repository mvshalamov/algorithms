package main

import (
	"fmt"
)

func buble_sort(input_array []int) []int {
	len_array := len(input_array) - 1
	for i := 0; i < len_array; i++ {
		swapped := false
		pass := len_array - i
		for j := 0; j < pass; j++ {
			if input_array[j] > input_array[j+1] {
				temp := input_array[j+1]
				input_array[j+1] = input_array[j]
				input_array[j] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return input_array
}

func main() {
	input_array := []int{1, 4, 3, 67, 2, 21, 3, 4, 325, 43, 5}
	fmt.Println(buble_sort(input_array))
}
