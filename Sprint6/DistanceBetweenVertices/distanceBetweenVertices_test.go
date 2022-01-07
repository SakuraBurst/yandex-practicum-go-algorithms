package distanceBetweenVertices_test

import (
	"bytes"
	"io"
	"log"
	"strings"
	"testing"

	distanceBetweenVertices "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/DistanceBetweenVertices"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestDistanceBetweenVertices(t *testing.T) {
	tests := []Test{{strings.NewReader(`5 5
2 4
3 5
2 1
2 3
4 5
1 5
`), "3"}, {strings.NewReader(`4 3
2 3
4 3
2 4
1 3
`), "-1"}, {strings.NewReader(`2 1
2 1
1 1
`), "0"}, {strings.NewReader(`7 6
2 3
6 7
1 3
4 5
7 5
6 5
6 4
`), "2"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			distanceBetweenVertices.DistanceBetweenVertices(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})
	}
}

func TestQueue_Push(t *testing.T) {
	queue := distanceBetweenVertices.Queue{
		CurrentLength: 16,
		QueueData:     make(distanceBetweenVertices.AdjacencyList, 16),
		PushTo:        0,
		TakeFrom:      0,
	}
	node := distanceBetweenVertices.Node{
		Value:    1,
		PointsTo: nil,
	}
	for i := 0; i < 1000000; i++ {
		queue.Push(&node)
	}
	log.Print(queue.CurrentLength)
}
