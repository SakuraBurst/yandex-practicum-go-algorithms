package partialSort_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	partialSort "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/PartialSort"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestPartialSort(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		partialSort.PartialSort(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}