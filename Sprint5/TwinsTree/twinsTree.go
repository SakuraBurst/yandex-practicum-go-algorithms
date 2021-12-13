package twinsTree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Solution(root1 *Node, root2 *Node) bool {
	// Your code
	// “ヽ(´▽｀)ノ”
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		if root1 != nil || root2 != nil {
			return false
		}
	}

	if Solution(root1.Left, root2.Left) && root1.Value == root2.Value && Solution(root1.Right, root2.Right) {
		return true
	}
	return false
}

func Solution2(root1 *Node, root2 *Node) bool {
	// Your code
	// “ヽ(´▽｀)ノ”
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		if root1 != nil || root2 != nil {
			return false
		}
	}
	leftTreeRes := Solution(root1.Left, root2.Left)

	if !leftTreeRes || root1.Value != root2.Value {
		return false
	}
	return Solution(root1.Right, root2.Right)
}
