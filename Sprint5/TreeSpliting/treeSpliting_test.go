package treeSpliting_test

import (
	"fmt"
	"testing"

	treeSpliting "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/TreeSpliting"
)

func TestTreeSpliting(t *testing.T) {
	node1 := &treeSpliting.Node{3, nil, nil, 1}
	node2 := &treeSpliting.Node{2, nil, node1, 2}
	node23 := &treeSpliting.Node{7, nil, nil, 1}
	node3 := &treeSpliting.Node{8, node23, nil, 2}
	node4 := &treeSpliting.Node{11, nil, nil, 1}
	node5 := &treeSpliting.Node{10, node3, node4, 4}
	node6 := &treeSpliting.Node{5, node2, node5, 7}
	left, right := treeSpliting.Split(node6, 1)
	if left.Size != 56 {
		printTree(left)
		t.Errorf("\nexpected 4 got %d", left.Size)
	}
	fmt.Println()
	if right.Size != 56 {
		printTree(right)
		t.Errorf("\nexpected 2 got %d", right.Size)
	}
}

func printTree(node *treeSpliting.Node) {
	if node == nil {
		return
	}
	printTree(node.Left)
	fmt.Print(node.Value)
	fmt.Print(" ")
	printTree(node.Right)
}
