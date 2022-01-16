package railroads

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Edge struct {
	weight int
	point  *Node
}

type PointsTo []Edge

type Node struct {
	value    int
	pointsTo PointsTo
}

type AdjacencyList []*Node

type PriorityQueue []Edge

func (q *PriorityQueue) Push(edge Edge) {
	*q = append(*q, edge)
	q.RebalanceTreeUp(q.Len() - 1)
}

func (q *PriorityQueue) Pop() Edge {
	// без defer)
	save := (*q)[1]
	(*q)[1] = (*q)[q.Len()-1]
	*q = (*q)[:q.Len()-1]
	return save
}

func (q PriorityQueue) RebalanceTreeUp(index int) {
	if index == 1 {
		return
	}
	parentIndex := index >> 1
	if q.Less(index, parentIndex) {
		q.Swap(index, parentIndex)
		q.RebalanceTreeUp(parentIndex)
	}
}

func (q PriorityQueue) RebalanceTreeDown(index int) {
	firstChild := index * 2
	secondChild := firstChild + 1
	if firstChild >= q.Len() {
		return
	}
	best := firstChild
	if secondChild < q.Len() && q.Less(secondChild, firstChild) {
		best = secondChild
	}
	if q.Less(best, index) {
		q.Swap(firstChild, index)
		q.RebalanceTreeDown(firstChild)
	}
}

func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Less(a, b int) bool {
	return q[a].weight > q[b].weight
}

func (q PriorityQueue) Swap(a, b int) {
	q[a], q[b] = q[b], q[a]
}

func Railroads(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	initData := strings.Fields(scanner.Text())
	vertexes, _ := strconv.Atoi(initData[0])
	edges, _ := strconv.Atoi(initData[1])
	AL := make(AdjacencyList, vertexes+1)
	for i := 0; i < vertexes; i++ {
		alValue := i + 1
		AL[alValue] = &Node{
			value:    alValue,
			pointsTo: nil,
		}
	}
	for i := 0; i < edges; i++ {
		scanner.Scan()
		edgeData := strings.Fields(scanner.Text())
		vertexA, _ := strconv.Atoi(edgeData[0])
		vertexB, _ := strconv.Atoi(edgeData[1])
		weight, _ := strconv.Atoi(edgeData[2])
		AL[vertexA].pointsTo = append(AL[vertexA].pointsTo, Edge{
			weight: weight,
			point:  AL[vertexB],
		})
		AL[vertexB].pointsTo = append(AL[vertexB].pointsTo, Edge{
			weight: weight,
			point:  AL[vertexA],
		})
	}
	//writer := bufio.NewWriter(w)
}
