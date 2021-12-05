package bulbs

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

var currentMax = 0
var isInit = true

func Bulbs(root *Node) int {
	if isInit {
		isInit = false
		currentMax = root.Value
	}
	if root.Value > currentMax {
		currentMax = root.Value
	}
	if root.Left != nil {
		Bulbs(root.Left)
	}
	if root.Right != nil {
		Bulbs(root.Right)
	}
	return currentMax
}
