package differentSerchTrees

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func DifferentSerchTrees(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	n, _ := strconv.Atoi(scaner.Text())
	writer := bufio.NewWriter(w)
	res := countAllUniqueBST(1, n)
	// fmt.Println(res)
	writer.WriteString(strconv.Itoa(len(res)))
	writer.Flush()
	// fmt.Println(reader)
}

func countAllUniqueBST(left, right int) []*Node {
	nodes := make([]*Node, 0)
	if left > right {
		// fmt.Println(left, right, nil)
		return append(nodes, nil)
	}
	for i := left; i <= right; i++ {
		leftNodes := countAllUniqueBST(left, i-1)
		rightNodes := countAllUniqueBST(i+1, right)
		// fmt.Println(i, left, right, leftNodes, rightNodes)
		for _, v := range leftNodes {
			ln := v
			for _, r := range rightNodes {
				rn := r

				nodes = append(nodes, &Node{Left: ln, Right: rn})
			}
		}
	}

	return nodes
}

func main() {
	DifferentSerchTrees(os.Stdin, os.Stdout)
}
