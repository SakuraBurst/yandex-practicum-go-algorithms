package allaround

type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}

func Solution(head *ListNode) *ListNode {
	if head.next == nil {
		swap(head)
		return head
	}
	result := Solution(head.next)
	swap(head)
	return result
}
func swap(node *ListNode) {
	node.next, node.prev = node.prev, node.next
}

func test() {
	node3 := ListNode{"node3", nil, nil}
	node2 := ListNode{"node2", &node3, nil}
	node1 := ListNode{"node1", &node2, nil}
	node0 := ListNode{"node0", &node1, nil}
	node3.prev = &node2
	node2.prev = &node1
	node1.prev = &node0
	Solution(&node0)
	// result is : newHead == node3
	// node3.next == node2
	// node2.next == node1
	// node2.prev = node3
	// node1.next == node0
	// node1.prev == node2
	// node0.prev == node1
}
