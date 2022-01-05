package maximumDistance

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value    int
	PointsTo AdjacencyList
}

type AdjacencyList []*Node

type Queue struct {
	CurrentSize  int
	Queue        AdjacencyList
	FirstInQueue int
	EmptyIndex   int
}

func (q *Queue) Push(node *Node) {
	q.Queue[q.EmptyIndex] = node
	q.EmptyIndex++
	q.EmptyIndex %= q.CurrentSize
	predictedIndex := (q.EmptyIndex + 1) % q.CurrentSize
	if q.FirstInQueue == predictedIndex {
		q.Rebalance()
	}
}

func (q *Queue) Pop() (*Node, error) {
	if q.FirstInQueue == q.EmptyIndex {
		return nil, errors.New("the queue is empty")
	}
	// тут я бы использовал defer, но это довольно нагруженная часть при ребалансировании, так что без него
	save := q.Queue[q.FirstInQueue]
	q.FirstInQueue++
	q.FirstInQueue %= q.CurrentSize
	return save, nil
}

func (q *Queue) Rebalance() {
	newQueue := make(AdjacencyList, q.CurrentSize*2)
	newQueueIndex := 0
	for i, err := q.Pop(); err == nil; i, err = q.Pop() {
		newQueue[newQueueIndex] = i
		newQueueIndex++
	}
	q.Queue = newQueue
	q.CurrentSize *= 2
	q.FirstInQueue = 0
	q.EmptyIndex = newQueueIndex
}

func MaximumDistance(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	initData := strings.Fields(scanner.Text())
	peaks, _ := strconv.Atoi(initData[0])
	edges, _ := strconv.Atoi(initData[1])
	AL := make(AdjacencyList, peaks+1)
	for i := 0; i < peaks; i++ {
		fixedIndex := i + 1
		AL[fixedIndex] = &Node{
			Value:    fixedIndex,
			PointsTo: nil,
		}
	}
	for i := 0; i < edges; i++ {
		scanner.Scan()
		edgeData := strings.Fields(scanner.Text())
		peakA, _ := strconv.Atoi(edgeData[0])
		peakB, _ := strconv.Atoi(edgeData[1])
		AL[peakA].PointsTo = append(AL[peakA].PointsTo, AL[peakB])
		AL[peakB].PointsTo = append(AL[peakB].PointsTo, AL[peakA])
	}
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	queue := Queue{
		CurrentSize:  16,
		Queue:        make(AdjacencyList, 16),
		FirstInQueue: 0,
		EmptyIndex:   0,
	}
	queue.Push(AL[N])
	colors := make([]string, peaks+1)
	distance := make([]int, peaks+1)
	max := bfs(queue, colors, distance, 0)
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(max))
	writer.Flush()
}

func bfs(queue Queue, colors []string, distance []int, currentMax int) int {
	node, err := queue.Pop()
	if err != nil {
		return currentMax
	}
	for _, n := range node.PointsTo {
		if colors[n.Value] == "" {
			colors[n.Value] = "gray"
			distance[n.Value] = distance[node.Value] + 1
			if distance[n.Value] > currentMax {
				currentMax = distance[n.Value]
			}
			queue.Push(n)
		}
	}
	colors[node.Value] = "black"
	return bfs(queue, colors, distance, currentMax)
}

func main() {
	MaximumDistance(os.Stdin, os.Stdout)
}
