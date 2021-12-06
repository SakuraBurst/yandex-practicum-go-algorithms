package searchTree_test

import (
	"testing"

	searchTree "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/SearchTree"
)

type Test struct {
	inputData  searchTree.Node
	outputData bool
}

func TestSearchTree(t *testing.T) {
	node1 := searchTree.Node{1, nil, nil}
	node2 := searchTree.Node{4, nil, nil}
	node2Wrong := searchTree.Node{5, nil, nil}
	node3 := searchTree.Node{3, &node1, &node2}
	node3Wrong := searchTree.Node{3, &node1, &node2Wrong}
	node69 := searchTree.Node{9, nil, nil}
	node4 := searchTree.Node{8, nil, &node69}
	trueNode := searchTree.Node{5, &node3, &node4}
	falseNode := searchTree.Node{5, &node3Wrong, &node4}
	tests := []Test{{trueNode, true}, {falseNode, false}}
	for _, v := range tests {
		res := searchTree.Solution(&v.inputData)
		if v.outputData != res {
			t.Errorf("\nexpected:\n%t\ngot:\n%t", v.outputData, res)
		}
	}
}
