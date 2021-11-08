package differenceOfTrashIndices_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	differenceOfTrashIndices "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/DifferenceOfTrashIndices"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestDifferenceOfTrashIndices(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		differenceOfTrashIndices.DifferenceOfTrashIndices(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}