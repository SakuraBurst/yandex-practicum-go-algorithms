package addNode_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	addNode "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/AddNode"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestAddNode(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		addNode.AddNode(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}