package differentSerchTrees_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	differentSerchTrees "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/DifferentSerchTrees"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestDifferentSerchTrees(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		differentSerchTrees.DifferentSerchTrees(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}