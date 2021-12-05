package treeSpliting_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	treeSpliting "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/TreeSpliting"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestTreeSpliting(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		treeSpliting.TreeSpliting(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}