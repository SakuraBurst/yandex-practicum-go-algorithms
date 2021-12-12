package filterUp_test

import (
	"testing"

	filterUp "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/FilterUp"
)

func TestFilterUp(t *testing.T) {
	sample := []int{-1, 12, 6, 8, 3, 15, 7}
	if filterUp.SiftUpRecurcial(sample, 5) != 1 {
		t.Error("WA")
	}

}
