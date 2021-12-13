package maxDepth_test

import (
	"testing"

	maxDepth "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/MaxDepth"
)

func TestMaxDepth(t *testing.T) {
	node1 := maxDepth.Node{1, nil, nil}
	node2 := maxDepth.Node{4, nil, nil}
	node3 := maxDepth.Node{3, &node1, &node2}
	node4 := maxDepth.Node{8, nil, nil}
	node5 := maxDepth.Node{5, &node3, &node4}
	r := maxDepth.Solution(&node5)
	if r != 3 {
		t.Errorf("\nexpected 3 got %d", r)
	}
}
