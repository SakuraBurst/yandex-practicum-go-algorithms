package addNode_test

import (
	"fmt"
	"testing"

	addNode "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/AddNode"
)

func TestAddNode(t *testing.T) {
	node1 := addNode.Node{7, nil, nil}
	node2 := addNode.Node{8, &node1, nil}
	node3 := addNode.Node{7, nil, &node2}
	newHead := addNode.Insert(&node3, 6)
	fmt.Println(newHead)
	if newHead != &node3 {
		t.Errorf("\nHEAD\nexpected:\n%v\ngot:\n%v", &node3, newHead)
	}
	if newHead.Left.Value != 6 {
		t.Errorf("\nValue\nexpected:\n%d\ngot:\n%d", node3.Value, 6)
	}

}
