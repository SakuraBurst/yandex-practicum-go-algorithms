package roadWeb

import (
	"bufio"
	"errors"
	"io"
	"os"
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

func (q *PriorityQueue) Pop() (Edge, error) {
	if q.Len() == 1 {
		return Edge{}, errors.New("queue is empty")
	}
	// без defer)
	save := (*q)[1]
	(*q)[1] = (*q)[q.Len()-1]
	*q = (*q)[:q.Len()-1]
	q.RebalanceTreeDown(1)
	return save, nil
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
		q.Swap(best, index)
		q.RebalanceTreeDown(best)
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

func RoadWeb(r io.Reader, w io.Writer) {
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

	res, visited := Dijkstra(AL)
	writer := bufio.NewWriter(w)
	if areAllNodesVisited(visited) {
		writer.WriteString(strconv.Itoa(res))
	} else {
		writer.WriteString("Oops! I did it again")
	}
	writer.Flush()
}

func Dijkstra(al AdjacencyList) (int, []bool) {
	queue := make(PriorityQueue, 1, cap(al))
	visited := make([]bool, cap(al))
	visited[1] = true
	accum := 0
	for _, edge := range al[1].pointsTo {
		queue.Push(edge)

	}
	for edge, err := queue.Pop(); err == nil; edge, err = queue.Pop() {

		if visited[edge.point.value] {
			continue
		}
		accum += edge.weight
		visited[edge.point.value] = true
		for _, e := range edge.point.pointsTo {
			if !visited[e.point.value] {
				queue.Push(e)
			}
		}
	}
	return accum, visited
}

func areAllNodesVisited(visited []bool) bool {
	for i := 1; i < len(visited); i++ {
		if !visited[i] {
			return false
		}
	}
	return true
}

func main() {
	RoadWeb(os.Stdin, os.Stdout)
}
