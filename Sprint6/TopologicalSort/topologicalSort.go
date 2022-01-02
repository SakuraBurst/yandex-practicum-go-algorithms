package topologicalSort

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value    int
	pointsTo AdjacencyList
}

type AdjacencyList []*Node

type Stack struct {
	next  int
	stack []string
}

func (s *Stack) String() string {
	return strings.Join(s.stack, " ")
}

func (s *Stack) Push(v string) {
	s.stack[s.next] = v
	s.next--
}

func TopologicalSort(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	initData := strings.Fields(scanner.Text())
	peaks, _ := strconv.Atoi(initData[0])
	edges, _ := strconv.Atoi(initData[1])
	AL := make(AdjacencyList, peaks+1)
	for i := 0; i < peaks; i++ {
		AL[i+1] = &Node{
			value:    i + 1,
			pointsTo: nil,
		}
	}
	for i := 0; i < edges; i++ {
		scanner.Scan()
		data := strings.Fields(scanner.Text())
		peakA, _ := strconv.Atoi(data[0])
		peakB, _ := strconv.Atoi(data[1])
		AL[peakA].pointsTo = append(AL[peakA].pointsTo, AL[peakB])
	}
	writer := bufio.NewWriter(w)
	colors := make([]string, peaks+1)
	stack := Stack{
		next:  peaks - 1,
		stack: make([]string, peaks),
	}
	for i := peaks; i > 0; i-- {
		if colors[i] == "" {
			TopoSort(AL[i], colors, &stack)
		}
	}
	_, err := writer.WriteString(fmt.Sprint(&stack))
	if err != nil {
		return
	}
	err = writer.Flush()
	if err != nil {
		return
	}
}

func TopoSort(n *Node, colors []string, stack *Stack) {
	colors[n.value] = "gray"
	for _, node := range n.pointsTo {
		if colors[node.value] == "" {
			TopoSort(node, colors, stack)
		}
	}
	colors[n.value] = "black"
	stack.Push(strconv.Itoa(n.value))
}

func main() {
	TopologicalSort(os.Stdin, os.Stdout)
}
