package dFS

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	pointsTo AdjacenceList
	value    int
}

type AdjacenceList []*Node

func DFS(r io.Reader, w io.Writer) {
	scaner := bufio.NewScanner(r)
	scaner.Scan()
	values := strings.Fields(scaner.Text())
	n, _ := strconv.Atoi(values[0])
	m, _ := strconv.Atoi(values[1])
	al := make(AdjacenceList, n+1)
	writer := bufio.NewWriter(w)
	for i := 0; i < m; i++ {
		scaner.Scan()
		edge := strings.Fields(scaner.Text())
		edgeA, _ := strconv.Atoi(edge[0])
		edgeB, _ := strconv.Atoi(edge[1])
		if al[edgeA] == nil {
			al[edgeA] = &Node{value: edgeA}
		}
		if al[edgeB] == nil {
			al[edgeB] = &Node{value: edgeB}
		}
		al[edgeA].pointsTo = append(al[edgeA].pointsTo, al[edgeB])
		al[edgeB].pointsTo = append(al[edgeB].pointsTo, al[edgeA])
	}
	scaner.Scan()

	start, _ := strconv.Atoi(scaner.Text())
	// 200ms
	// var dfcRecurcial func(node *Node) error
	// colors := make([]string, n+1)
	// dfcRecurcial = func(node *Node) error {
	// 	if node == nil {
	// 		return errors.New("node doesen't exist")
	// 	}
	// 	colors[node.value] = "black"
	// 	writer.WriteString(strconv.Itoa(node.value) + " ")
	// 	node.pointsTo = mergeSrot(node.pointsTo)
	// 	for _, v := range node.pointsTo {
	// 		if colors[v.value] != "black" {
	// 			dfcRecurcial(v)
	// 		}

	// 	}
	// 	return nil

	// }
	// if err := dfcRecurcial(al[start]); err != nil {
	// 	writer.WriteString(strconv.Itoa(start))
	// }

	// 196
	if err := dfcRecurcialV2(al[start], make([]string, n+1), writer); err != nil {
		writer.WriteString(strconv.Itoa(start))
	}

	// 1 second+; alloc?
	// res := dfcRecurcial(al[start], make([]string, n+1), nil)
	// for i := len(res) - 1; i >= 0; i-- {
	// 	writer.WriteString(res[i] + " ")
	// }

	//248ms
	// res, err := dfcMin(al[start], make([]string, n+1))

	// if err != nil {
	// 	writer.WriteString(strconv.Itoa(start))
	// } else {
	// 	writer.WriteString(strings.Join(res, " "))
	// }

	writer.Flush()
	//

}

type Stack []*Node

func (s *Stack) Push(i *Node) {
	*s = append(*s, i)
}

func (s *Stack) Pop() (*Node, error) {
	if len(*s) == 0 {
		return nil, errors.New("stack is empty")
	}
	defer func() {
		*s = (*s)[0 : len(*s)-1]
	}()
	return (*s)[len(*s)-1], nil
}

func (s Stack) Len() int {
	return len(s)
}

func dfc(node *Node, colors []string) ([]string, error) {
	if node == nil {
		return nil, errors.New("no such node")
	}

	stack := make(Stack, 0)

	stack.Push(node)

	result := make([]string, 0)

	for stack.Len() != 0 {

		n, _ := stack.Pop()

		if colors[n.value] == "" {
			colors[n.value] = "gray"
			stack.Push(n)
			result = append(result, strconv.Itoa(n.value))
			n.pointsTo = mergeSrot(n.pointsTo)
			for _, v := range n.pointsTo {
				if colors[v.value] == "" {
					stack.Push(v)
				}
			}
		} else if colors[n.value] == "gray" {
			colors[n.value] = "black"

		}
	}
	return result, nil
}

func dfcMin(node *Node, colors []string) ([]string, error) {
	if node == nil {
		return nil, errors.New("no such node")
	}
	stack := make(Stack, 0)
	stack.Push(node)
	result := make([]string, 0)
	for stack.Len() != 0 {
		n, _ := stack.Pop()
		if colors[n.value] == "" {
			colors[n.value] = "black"
			result = append(result, strconv.Itoa(n.value))
			n.pointsTo = mergeSrot(n.pointsTo)
			for _, v := range n.pointsTo {
				if colors[v.value] == "" {
					stack.Push(v)
				}
			}
		}
	}
	return result, nil
}

func mergeSrot(l AdjacenceList) AdjacenceList {
	if len(l) <= 1 {
		return l
	}
	mid := len(l) >> 1
	a := mergeSrot(l[:mid])
	b := mergeSrot(l[mid:])
	res := make(AdjacenceList, 0, len(a)+len(b))
	aIndex := 0
	bIndex := 0
	for aIndex < len(a) || bIndex < len(b) {
		if aIndex == len(a) {
			res = append(res, b[bIndex])
			bIndex++
			continue
		}
		if bIndex == len(b) {
			res = append(res, a[aIndex])
			aIndex++
			continue
		}
		// <= if recursial, > if stack
		if a[aIndex].value > b[bIndex].value {
			res = append(res, a[aIndex])
			aIndex++
		} else {
			res = append(res, b[bIndex])
			bIndex++
		}
	}
	return res
}

func main() {
	DFS(os.Stdin, os.Stdout)
}

func dfcRecurcial(node *Node, colors []string, result []string) []string {
	if colors[node.value] == "black" {
		return nil
	}
	colors[node.value] = "black"

	for _, v := range node.pointsTo {
		if colors[v.value] != "black" {
			result = append(result, dfcRecurcial(v, colors, nil)...)
		}

	}
	result = append(result, strconv.Itoa(node.value))
	return result
}

func dfcRecurcialV2(node *Node, colors []string, wr io.StringWriter) error {
	if node == nil {
		return errors.New("node doesen't exist")
	}

	if colors[node.value] == "black" {
		return nil
	}
	colors[node.value] = "black"
	wr.WriteString(strconv.Itoa(node.value) + " ")
	// node.pointsTo = mergeSrot(node.pointsTo)
	for _, v := range node.pointsTo {
		if colors[v.value] != "black" {
			dfcRecurcialV2(v, colors, wr)
		}

	}
	return nil
}
