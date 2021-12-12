package filterDown_test

import (
	"testing"

	filterDown "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/FilterDown"
)

func TestFilterDown(t *testing.T) {
	sample := []int{-1, 12, 1, 8, 3, 4, 7}
	if filterDown.SiftDown2Recurcial(sample, 2) != 5 {
		t.Error("WA")
	}
}
