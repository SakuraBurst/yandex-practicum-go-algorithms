package balancedTree_test

import (
	"testing"

	balancedTree "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/BalancedTree"
)

func TestBalancedTree(t *testing.T) {
	node1 := balancedTree.Node{1, nil, nil}
	node2 := balancedTree.Node{-5, nil, nil}
	node3 := balancedTree.Node{3, &node1, &node2}
	node4 := balancedTree.Node{10, nil, nil}
	node5 := balancedTree.Node{2, &node3, &node4}
	if !balancedTree.Solution(&node5) {
		t.Errorf("\nexpected:\n%t\ngot:\n%t", true, false)
	}

	node28 := balancedTree.Node{8, nil, nil}
	node27 := balancedTree.Node{7, nil, nil}
	node26 := balancedTree.Node{6, nil, nil}
	node25 := balancedTree.Node{5, nil, nil}
	node24 := balancedTree.Node{4, &node27, &node28}
	node23 := balancedTree.Node{1, &node25, &node26}
	node22 := balancedTree.Node{1, nil, &node24}
	node21 := balancedTree.Node{1, &node23, nil}
	node20 := balancedTree.Node{1, &node21, &node22}
	if balancedTree.Solution(&node20) {
		t.Errorf("\nexpected:\n%t\ngot:\n%t", false, true)
	}
}

func test() {

}
