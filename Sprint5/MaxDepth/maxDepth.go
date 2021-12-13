package maxDepth

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Solution(root *Node) int {
	// Your code
	// “ヽ(´▽｀)ノ”
	if root == nil {
		return 0
	}

	l := Solution(root.Left)
	r := Solution(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}
