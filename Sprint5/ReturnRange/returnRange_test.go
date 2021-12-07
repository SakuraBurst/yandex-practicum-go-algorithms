package returnRange_test

import (
	"fmt"
	"testing"

	returnRange "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/ReturnRange"
)

func TestReturnRange(t *testing.T) {
	node11 := returnRange.Node{2, nil, nil}
	node12 := returnRange.Node{1, nil, &node11}
	node13 := returnRange.Node{8, nil, nil}
	node14 := returnRange.Node{8, nil, &node13}
	node15 := returnRange.Node{9, &node14, nil}
	node16 := returnRange.Node{10, &node15, nil}
	node17 := returnRange.Node{5, &node12, &node16}

	node21 := returnRange.Node{15, nil, nil}
	node23 := returnRange.Node{1, nil, nil}
	node22 := returnRange.Node{3, &node23, &node11}
	node24 := returnRange.Node{11, nil, &node21}
	node25 := returnRange.Node{7, &node22, &node24}

	returnRange.PrintRange(&node17, 2, 8)
	fmt.Println()
	returnRange.PrintRange(&node25, 6, 14)
	fmt.Println()
	t.Fatal(t.Name())
	// expected output: 2 5 8 8

}
