package addNode

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Insert(root *Node, key int) *Node {
	// Your code
	// “ヽ(´▽｀)ノ”
	if root.Value <= key {
		if root.Right == nil {
			root.Right = &Node{Value: key}
		} else {
			Insert(root.Right, key)
		}
		return root
	} else {
		if root.Left == nil {
			root.Left = &Node{Value: key}
		} else {
			Insert(root.Left, key)
		}
		return root
	}
}
