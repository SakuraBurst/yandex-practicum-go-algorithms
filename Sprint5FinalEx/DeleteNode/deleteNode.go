package deleteNode

import "fmt"

// для тестирования
type TestNode struct {
	Value int
	Left  *TestNode
	Right *TestNode
}

// для тестирования
func Remove(node *TestNode, key int) *TestNode {
	if node == nil {
		return node
	}
	if node.Value == key {
		if node.Left == nil && node.Right == nil {
			return nil
		}

		if node.Left == nil {
			return node.Right
		}

		if node.Right == nil {
			return node.Left
		}

		val, replaceNode := getAndDeleteSmallestValInTree(node.Right)
		fmt.Println(val)
		node.Right = replaceNode
		node.Value = val
		return node
	}
	if node.Value > key {
		newNode := Remove(node.Left, key)
		node.Left = newNode
	} else {
		newNode := Remove(node.Right, key)
		node.Right = newNode
	}

	return node
}

func getAndDeleteSmallestValInTree(n *TestNode) (int, *TestNode) {
	if n.Left == nil {
		if n.Right == nil {
			return n.Value, nil
		}
		return n.Value, n.Left
	}
	val, replaceNode := getAndDeleteSmallestValInTree(n.Left)
	n.Left = replaceNode
	return val, n
}
