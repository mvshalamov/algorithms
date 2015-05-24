package main

import (
	"fmt"
)

func main() {
	a := []int{10, 2, 34, 1, 23, 3, 56, 3}
	bin := binary_heap{heap: a}
	fmt.Println(a)
	fmt.Println(bin.heap_sort(a))
}

type binary_heap struct {
	heap []int
}

func (h *binary_heap) size() int {
	return len(h.heap) - 1
}

func (h *binary_heap) heapify(root int) {
	for true {
		left_ch := 2*root + 1
		right_ch := 2*root + 2
		largest := root

		if left_ch <= h.size() && h.heap[left_ch] > h.heap[largest] {
			largest = left_ch
		}

		if right_ch <= h.size() && h.heap[right_ch] > h.heap[largest] {
			largest = right_ch
		}

		if largest == root {
			break
		}

		it := h.heap[largest]
		h.heap[largest] = h.heap[root]
		h.heap[root] = it
		root = largest
	}
}

func (h *binary_heap) built_heap(data []int) {
	h.heap = data
	num := h.size() / 2
	for i := num; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *binary_heap) get_max() int {
	res := h.heap[0]
	h.heap[0] = h.heap[h.size()]
	h.heap = h.heap[0:h.size()]
	h.heapify(0)
	return res
}

func (h *binary_heap) heap_sort(data []int) []int {
	h.built_heap(data)
	res := make([]int, len(data))
	num := h.size()
	for i := 0; i <= num; i++ {

		res[i] = h.get_max()
	}
	return res
}
