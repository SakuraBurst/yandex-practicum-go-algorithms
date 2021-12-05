package maxPathInTree_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	maxPathInTree "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/MaxPathInTree"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestMaxPathInTree(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		maxPathInTree.MaxPathInTree(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}