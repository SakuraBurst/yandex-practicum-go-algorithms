package treeAnagramm_test

import (
	"testing"

	treeAnagramm "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/TreeAnagramm"
)

func TestTreeAnagramm(t *testing.T) {
	node1 := treeAnagramm.Node{3, nil, nil}
	node2 := treeAnagramm.Node{4, nil, nil}
	node3 := treeAnagramm.Node{4, nil, nil}
	node4 := treeAnagramm.Node{3, nil, nil}
	node5 := treeAnagramm.Node{2, &node1, &node2}
	node6 := treeAnagramm.Node{2, &node3, &node4}
	node7 := treeAnagramm.Node{1, &node5, &node6}

	if !treeAnagramm.Solution(&node7) {
		t.Error("WA")
	}
	node11 := treeAnagramm.Node{0, nil, nil}
	node12 := treeAnagramm.Node{1, nil, nil}
	node13 := treeAnagramm.Node{1, nil, nil}
	node14 := treeAnagramm.Node{0, nil, nil}
	node15 := treeAnagramm.Node{3, &node12, &node11}
	node16 := treeAnagramm.Node{3, &node14, &node13}
	node17 := treeAnagramm.Node{2, nil, &node15}
	node18 := treeAnagramm.Node{2, &node16, nil}
	node19 := treeAnagramm.Node{0, &node18, &node17}
	if !treeAnagramm.Solution(&node19) {
		t.Error("WA")
	}
}
