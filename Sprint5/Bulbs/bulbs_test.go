package bulbs_test

import (
	"testing"

	bulbs "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/Bulbs"
)

type Test struct {
	inputData  bulbs.Node
	outputData int
}

func TestBulbs(t *testing.T) {
	node1 := bulbs.Node{1, nil, nil}
	node2 := bulbs.Node{-5, nil, nil}
	node3 := bulbs.Node{3, &node1, &node2}
	node4 := bulbs.Node{2, &node3, nil}
	tests := []Test{{node4, 3}}
	for _, v := range tests {
		res := bulbs.Bulbs(&v.inputData)
		if res != v.outputData {
			t.Errorf("\nexpected:\n%d\ngot:\n%d", v.outputData, res)
		}
	}
}
