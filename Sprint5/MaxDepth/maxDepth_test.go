package maxDepth_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	maxDepth "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/MaxDepth"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestMaxDepth(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		maxDepth.MaxDepth(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}