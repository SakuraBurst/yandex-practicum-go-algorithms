package worstjob

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		return head.next
	}
	var savedNode *ListNode
	var currentNode = head
	count := 0
	for count < idx {
		if count == idx-1 {
			savedNode = currentNode
		}
		currentNode = currentNode.next
		count++

	}
	savedNode.next = currentNode.next
	return head
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	/*newHead :=*/ Solution(&node0, 1)
	// result is : node0 -> node2 -> node3
}
