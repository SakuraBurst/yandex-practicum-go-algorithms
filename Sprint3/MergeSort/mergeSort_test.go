package mergeSort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	a := []int{1, 4, 9, 2, 10, 11}
	b := merge(a, 0, 3, 6)
	expected := []int{1, 2, 4, 9, 10, 11}
	if !reflect.DeepEqual(b, expected) {
		panic("WA. Merge")
	}

	c := []int{1, 4, 2, 10, 1, 2}
	merge_sort(c, 0, 6)
	expected = []int{1, 1, 2, 2, 4, 10}
	if !reflect.DeepEqual(c, expected) {
		panic("WA. MergeSort")
	}

	bigArray := make([]int, 50001)
	merge_sort(bigArray, 0, 50000)
}
