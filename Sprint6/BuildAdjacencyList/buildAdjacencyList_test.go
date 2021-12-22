package buildAdjacencyList_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	buildAdjacencyList "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/BuildAdjacencyList"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBuildAdjacencyList(t *testing.T) {
	tests := []Test{{strings.NewReader(`5 3
1 3
2 3
5 2
`), `1 3 
1 3 
0 
0 
1 2 
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		buildAdjacencyList.BuildAdjacencyList(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
