package topologicalSort_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	topologicalSort "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/TopologicalSort"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestTopologicalSort(t *testing.T) {
	tests := []Test{{strings.NewReader(`5 3
3 2
3 4
2 5
`), "1 3 2 4 5"}, {strings.NewReader(`6 3
6 4
4 1
5 1
`), "2 3 5 6 4 1"}, {strings.NewReader(`4 0
`), "1 2 3 4"}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		topologicalSort.TopologicalSort(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
