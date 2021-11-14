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

func TestPartialSort(t *testing.T) {
	tests := []Test{{strings.NewReader(`4
0 1 3 2`), "3"}, {strings.NewReader(`8
3 6 7 4 1 5 0 2`), "1"}, {strings.NewReader(`5
1 0 2 3 4`), "4"}, {strings.NewReader(`4
2 1 0 3`), "2"}, {strings.NewReader(`10
2 7 0 3 6 4 1 5 8 9`), "3"}, {strings.NewReader(`5
3 1 2 0 4`), "2"}, {strings.NewReader(`3
2 0 1`), "1"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		partialSort.PartialSort(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
