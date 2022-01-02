package convertListOfEdgesToAdjacencyMatrix_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	convertListOfEdgesToAdjacencyMatrix "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/ConvertListOfEdgesToAdjacencyMatrix"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestConvertListOfEdgesToAdjacencyMatrix(t *testing.T) {
	tests := []Test{{strings.NewReader(`5 3
1 3
2 3
5 2
`), `0 0 1 0 0 
0 0 1 0 0 
0 0 0 0 0 
0 0 0 0 0 
0 1 0 0 0 
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		convertListOfEdgesToAdjacencyMatrix.ConvertListOfEdgesToAdjacencyMatrix(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}

func BenchmarkConvertListOfEdgesToAdjacencyMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
