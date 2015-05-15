package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	data := []int{6, 1, 7, 2, 9}
	res := []int{1, 2, 6, 7, 9}
	quicksort(data, 0, len(data)-1)
	if testEq(data, res) == false {
		t.Error("not work")
	}
}

func testEq(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
