package bFS_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	bFS "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/BFS"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBFS(t *testing.T) {
	tests := []Test{{strings.NewReader(`4 4
1 2
2 3
3 4
1 4
3
`), "3 2 4 1 "}, {strings.NewReader(`2 1
2 1
1
`), "1 2 "}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		bFS.BFS(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}

func TestQueue_Push(t *testing.T) {
	testNode := &bFS.Node{
		Value:    1,
		PointsTo: nil,
	}
	queue := bFS.Queue{
		CurrentSize:  16,
		Queue:        make(bFS.AdjacencyList, 16),
		FirstInQueue: 0,
		EmptyIndex:   0,
	}
	for i := 0; i < 1000; i++ {
		queue.Push(testNode)
	}
}
