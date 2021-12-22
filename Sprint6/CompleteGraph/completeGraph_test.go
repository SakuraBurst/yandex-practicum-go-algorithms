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

func TestCompleteGraph(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		completeGraph.CompleteGraph(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}