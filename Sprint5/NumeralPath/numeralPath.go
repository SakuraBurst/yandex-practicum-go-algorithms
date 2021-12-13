package numeralPath

import "strconv"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Solution(root *Node) int {
	// Your code
	// “ヽ(´▽｀)ノ”
	return SolutionCheat(root, "")
}

func SolutionCheat(root *Node, s string) int {
	currentIntS := s + strconv.Itoa(root.Value)
	currentInt, _ := strconv.Atoi(currentIntS)
	if root.Left == nil && root.Right == nil {
		return currentInt
	}
	if root.Left != nil && root.Right != nil {
		return SolutionCheat(root.Left, currentIntS) + SolutionCheat(root.Right, currentIntS)
	}
	if root.Left != nil {
		return SolutionCheat(root.Left, currentIntS)
	}
	return SolutionCheat(root.Right, currentIntS)
}
