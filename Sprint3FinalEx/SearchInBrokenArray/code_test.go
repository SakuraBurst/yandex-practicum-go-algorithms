package brokenarray_test

import (
	"fmt"
	"testing"

	brokenarray "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3FinalEx/SearchInBrokenArray"
)

func Test(t *testing.T) {
	//            0   1    2    3    4    5    6    7    8    9   10  11 12 13 14 15
	arr := []int{19, 21, 100, 101, 109, 112, 113, 114, 115, 116, 153, 1, 4, 5, 7, 12}
	// arr = []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if brokenarray.BrokenSearch(arr, 5) != 13 {
		fmt.Println(brokenarray.BrokenSearch(arr, 5), "kachalka")
		panic("WA")
	}
}
