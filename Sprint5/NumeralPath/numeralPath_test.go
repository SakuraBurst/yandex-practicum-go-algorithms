package numeralPath_test

import (
	"testing"

	numeralPath "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/NumeralPath"
)

func TestNumeralPath(t *testing.T) {
	node1 := numeralPath.Node{2, nil, nil}
	node2 := numeralPath.Node{1, nil, nil}
	node3 := numeralPath.Node{3, &node1, &node2}
	node4 := numeralPath.Node{2, nil, nil}
	node5 := numeralPath.Node{1, &node4, &node3}
	r := numeralPath.Solution(&node5)
	if r != 275 {
		t.Errorf("\nexpected 275 got %d", r)
	}
}
