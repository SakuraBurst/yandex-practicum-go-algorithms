package treeSpliting

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
	Size  int
}

func Split(node *Node, k int) (*Node, *Node) {
	// Your code
	// “ヽ(´▽｀)ノ”
	if node == nil {
		return nil, nil
	}

	leftSize, rightSize := subTreesLength(node.Left, node.Right)
	fmt.Println("k = ", k, "current value = ", node.Value, "leftSize = ", leftSize, "right size = ", rightSize)
	if leftSize+1 > k {
		ln, rn := Split(node.Left, k)
		_, rnSize := subTreesLength(ln, rn)

		node.Size = node.Size - leftSize + rnSize
		node.Left = rn
		return ln, node
	}
	ln, rn := Split(node.Right, k-(leftSize+1))
	lnSize, _ := subTreesLength(ln, rn)
	node.Size = node.Size - rightSize + lnSize
	node.Right = ln
	return node, rn
}

func subTreesLength(l, r *Node) (leftSize, rightSize int) {
	if l != nil {
		leftSize = l.Size
	}
	if r != nil {
		rightSize = r.Size
	}
	return
}
