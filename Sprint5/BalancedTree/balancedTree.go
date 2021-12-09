package balancedTree

import (
	"errors"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Solution(root *Node) bool {
	// Your code
	// “ヽ(´▽｀)ノ”
	if _, err := counter(root); err != nil {
		return false
	}
	return true
}

func counter(root *Node) (int, error) {
	if root == nil {
		return 0, nil
	}

	leftTreeLength, leftErr := counter(root.Left)
	rightTreeLength, rightErr := counter(root.Right)
	if leftErr != nil || rightErr != nil {
		return 0, errors.New("tree is not balanced")
	}
	res := leftTreeLength - rightTreeLength
	if res > 1 || res < -1 {
		return 0, errors.New("tree is not balanced")
	}
	if leftTreeLength >= rightTreeLength {
		return leftTreeLength + 1, nil
	} else {
		return rightTreeLength + 1, nil
	}

}
