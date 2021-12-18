package deleteNode_test

import (
	"testing"

	deleteNode "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5FinalEx/DeleteNode"
)

func TestDeleteNode(t *testing.T) {
	node1 := deleteNode.TestNode{2, nil, nil}
	node2 := deleteNode.TestNode{3, &node1, nil}
	node3 := deleteNode.TestNode{1, nil, &node2}
	node4 := deleteNode.TestNode{6, nil, nil}
	node5 := deleteNode.TestNode{8, &node4, nil}
	node6 := deleteNode.TestNode{10, &node5, nil}
	node7 := deleteNode.TestNode{5, &node3, &node6}
	newHead := deleteNode.Remove(&node7, 10)
	if newHead.Value != 5 {
		t.Errorf("\nnewHead.Value: expect 5 got %d", newHead.Value)
	}
	if newHead.Right != &node5 {
		t.Errorf("\nnewHead.Right: expect %v got %v", &node5, newHead.Right)
	}
	if newHead.Right.Value != 8 {
		t.Errorf("\nnewHead.Right.Value: expect 8 got %d", newHead.Right.Value)
	}

	node1_1 := deleteNode.TestNode{2, nil, nil}
	node1_2 := deleteNode.TestNode{3, &node1_1, nil}
	node1_3 := deleteNode.TestNode{1, nil, &node1_2}
	node1_4 := deleteNode.TestNode{6, nil, nil}
	node1_n1 := deleteNode.TestNode{11, nil, nil}
	node1_n2 := deleteNode.TestNode{16, nil, nil}
	node1_n3 := deleteNode.TestNode{15, &node1_n1, &node1_n2}
	node1_5 := deleteNode.TestNode{8, &node1_4, nil}
	node1_6 := deleteNode.TestNode{10, &node1_5, &node1_n3}
	node1_7 := deleteNode.TestNode{5, &node1_3, &node1_6}
	newHead = deleteNode.Remove(&node1_7, 10)
	if newHead.Value != 5 {
		t.Errorf("\nnewHead.Value: expect 5 got %d", newHead.Value)
	}
	if newHead.Right.Value != 11 {
		t.Errorf("\nnewHead.Right.Value: expect 11 got %d", newHead.Right.Value)
	}

}
