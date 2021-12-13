package maxPathInTree_test

import (
	"testing"

	maxPathInTree "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/MaxPathInTree"
)

func TestMaxPathInTree(t *testing.T) {
	node1 := maxPathInTree.Node{5, nil, nil}
	node2 := maxPathInTree.Node{1, nil, nil}
	node3 := maxPathInTree.Node{-3, &node2, &node1}
	node4 := maxPathInTree.Node{2, nil, nil}
	node5 := maxPathInTree.Node{2, &node4, &node3}
	r := maxPathInTree.Solution(&node5)
	if r != 6 {
		t.Errorf("\nexpected 3 got %d", r)
	}
}
