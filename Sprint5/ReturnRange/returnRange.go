package returnRange

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func PrintRange(root *Node, left int, right int) {
	if root.Left == nil && root.Right == nil {
		if root.Value >= left && root.Value <= right {
			fmt.Print(root.Value)
			fmt.Print(" ")
		}
		return
	}
	if root.Left != nil && root.Value >= left {
		PrintRange(root.Left, left, right)
		if root.Value <= right {
			fmt.Print(root.Value)
			fmt.Print(" ")
		}

	}
	if root.Right != nil && root.Value <= right {
		if root.Value >= left && root.Left == nil {
			fmt.Print(root.Value)
			fmt.Print(" ")
		}
		PrintRange(root.Right, left, right)

	}
	// Your code
	// “ヽ(´▽｀)ノ”
}
