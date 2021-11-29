package sumOfFour_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	sumOfFour "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/SumOfFour"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestSumOfFour(t *testing.T) {
	tests := []Test{{strings.NewReader(`8
10
2 3 2 4 1 10 3 0
`), `3
0 3 3 4
1 2 3 4
2 2 3 3
`}, {strings.NewReader(`6
0
1 0 -1 0 2 -2
`), `3
-2 -1 1 2
-2 0 0 2
-1 0 0 1
`}, {strings.NewReader(`5
4
1 1 1 1 1
`), `1
1 1 1 1
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		sumOfFour.SumOfFour(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
