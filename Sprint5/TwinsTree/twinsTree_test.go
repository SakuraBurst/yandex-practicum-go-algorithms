package twinsTree_test

import (
	"testing"

	twinsTree "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/TwinsTree"
)

func TestTwinsTree(t *testing.T) {
	node1 := twinsTree.Node{1, nil, nil}
	node2 := twinsTree.Node{2, nil, nil}
	node3 := twinsTree.Node{3, &node1, &node2}
	node4 := twinsTree.Node{1, nil, nil}
	node5 := twinsTree.Node{2, nil, nil}
	node6 := twinsTree.Node{3, &node4, &node5}

	if !twinsTree.Solution2(&node3, &node6) {
		t.Error("WA")
	}
}
