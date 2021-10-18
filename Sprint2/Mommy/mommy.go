package mommy

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, elem string) int {
	index := 0
	for head != nil && head.data != elem {
		index++
		head = head.next
	}
	if head == nil {
		return -1
	} else {
		return index
	}
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	/*idx :=*/ Solution(&node0, "node2")
	// result is : idx == 2
}
