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

func TestTopologicalSort(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		topologicalSort.TopologicalSort(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}