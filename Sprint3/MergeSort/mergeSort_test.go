package mergeSort_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	mergeSort "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/MergeSort"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestMergeSort(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		mergeSort.MergeSort(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}