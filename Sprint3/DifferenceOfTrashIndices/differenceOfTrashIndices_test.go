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

func TestDifferenceOfTrashIndices(t *testing.T) {
	tests := []Test{{strings.NewReader(`3
2 3 4
2`), "1"}, {strings.NewReader(`3
1 3 1
1`), "0"}, {strings.NewReader(`3
1 3 5
3`), "4"}, {strings.NewReader(`4
9 6 10 3
3`), "3"}, {strings.NewReader(`6
27 67 3 4 17 98
7`), "31"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		differenceOfTrashIndices.DifferenceOfTrashIndices(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
