package completeGraph

import (
	"bufio"
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

func CompleteGraph(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	initData := strings.Fields(scanner.Text())
	vertexes, _ := strconv.Atoi(initData[0])
	edges, _ := strconv.Atoi(initData[1])
	AL := make(AdjacencyList, vertexes+1)
	for i := 0; i < vertexes; i++ {
		correctIndex := i + 1
		AL[correctIndex] = &Node{
			value:    correctIndex,
			pointsTo: nil,
		}
	}
	for i := 0; i < edges; i++ {
		scanner.Scan()
		edgeData := strings.Fields(scanner.Text())
		vertexA, _ := strconv.Atoi(edgeData[0])
		vertexB, _ := strconv.Atoi(edgeData[1])
		if vertexA != vertexB {
			AL[vertexA].pointsTo = append(AL[vertexA].pointsTo, AL[vertexB])
			AL[vertexB].pointsTo = append(AL[vertexB].pointsTo, AL[vertexA])
		}
	}
	writer := bufio.NewWriter(w)
External:
	for _, node := range AL {
		if node == nil {
			continue
		}
		visited := make([]bool, vertexes+1)
		for _, n := range node.pointsTo {
			visited[n.value] = true
		}
		for i, isVisited := range visited {
			if i == 0 || i == node.value {
				continue
			}
			if !isVisited {
				writer.WriteString("NO")
				break External
			}
		}
	}
	s := writer.Buffered()
	if s == 0 {
		writer.WriteString("YES")
	}
	writer.Flush()
}

func main() {
	CompleteGraph(os.Stdin, os.Stdout)
}
