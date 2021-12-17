package deleteNode_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	deleteNode "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5FinalEx/DeleteNode"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestDeleteNode(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		deleteNode.DeleteNode(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}