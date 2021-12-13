package treeAnagramm

import (
	"fmt"
	"strconv"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Solution(root *Node) bool {
	// Your code
	// “ヽ(´▽｀)ノ”
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return false
	}
	l := LeftMidRightValue(root.Left)
	r := RightMidLeftValue(root.Right)
	fmt.Println(l, r)
	return LeftMidRightValue(root.Left) == RightMidLeftValue(root.Right)
}

func LeftMidRightValue(root *Node) string {
	if root == nil {
		return ""
	}
	return LeftMidRightValue(root.Left) + strconv.Itoa(root.Value) + LeftMidRightValue(root.Right)
}

func RightMidLeftValue(root *Node) string {
	if root == nil {
		return ""
	}
	return RightMidLeftValue(root.Right) + strconv.Itoa(root.Value) + RightMidLeftValue(root.Left)
}
