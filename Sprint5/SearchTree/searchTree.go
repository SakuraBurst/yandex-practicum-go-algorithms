package searchTree

import "errors"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Solution(root *Node) bool {
	// Your code
	// “ヽ(´▽｀)ノ”
	_, err := isSearchTree(root)
	if err != nil {
		return false
	}
	return true
}

func isSearchTree(node *Node) ([]int, error) {
	if node == nil {
		return []int{}, nil
	}
	if node.Right == nil && node.Left == nil {
		return []int{node.Value}, nil
	}
	leftValues, leftErr := isSearchTree(node.Left)
	rightValues, rightErr := isSearchTree(node.Right)
	if leftErr != nil || rightErr != nil {
		return nil, errors.New("wrong tree")
	}
	for _, v := range leftValues {
		if v >= node.Value {
			return nil, errors.New("wrong tree")
		}
	}
	for _, v := range rightValues {
		if v <= node.Value {
			return nil, errors.New("wrong tree")
		}
	}
	return append(leftValues, rightValues...), nil
}
