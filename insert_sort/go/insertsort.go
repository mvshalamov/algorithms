package main

import (
	"fmt"
)

func main() {
	data := []int{4, 2, 5, 1, 7, 3, 5, 9}
	fmt.Println(data)
	fmt.Println(insertsort(data))
	fmt.Println(data)
}

func insertsort(data []int) []int {
	for pos, val := range data {
		i := pos
		for i > 0 && data[i-1] > val {
			data[i] = data[i-1]
			i--
		}

		data[i] = val
	}

	return data
}
