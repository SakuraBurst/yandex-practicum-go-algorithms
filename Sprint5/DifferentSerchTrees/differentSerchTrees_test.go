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

func TestDifferentSerchTrees(t *testing.T) {
	tests := []Test{{strings.NewReader("2"), "2"}, {strings.NewReader("3"), "5"}, {strings.NewReader("4"), "14"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		differentSerchTrees.DifferentSerchTrees(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
