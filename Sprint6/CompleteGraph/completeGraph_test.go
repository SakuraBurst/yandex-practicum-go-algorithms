package completeGraph_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	completeGraph "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/CompleteGraph"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestCompleteGraph(t *testing.T) {
	tests := []Test{{strings.NewReader(`4 6
1 2
2 2
2 3
2 4
3 4
4 3
`), "NO"}, {strings.NewReader(`3 5
1 2
2 1
3 1
2 3
3 3
`), "YES"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			completeGraph.CompleteGraph(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})
	}
}
