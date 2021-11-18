package brokenarray_test

import (
	"fmt"
	"testing"

	brokenarray "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3FinalEx/SearchInBrokenArray"
)

func Test(t *testing.T) {
	arr := []int{19, 21, 100, 101, 109, 112, 113, 114, 115, 116, 153, 1, 4, 5, 7, 12}
	if brokenarray.BrokenSearch(arr, 5) != 6 {
		fmt.Println(brokenarray.BrokenSearch(arr, 5), "kachalka")
		panic("WA")
	}
}
