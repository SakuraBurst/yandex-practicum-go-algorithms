package maximumDistance_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	maximumDistance "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/MaximumDistance"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestMaximumDistance(t *testing.T) {
	tests := []Test{{strings.NewReader(`5 4
2 1
4 5
4 3
3 2
2
`), "3"}, {strings.NewReader(`3 3
3 1
1 2
2 3
1
`), "1"}, {strings.NewReader(`6 8
6 1
1 3
5 1
3 5
3 4
6 5
5 2
6 2
4
`), "3"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			maximumDistance.MaximumDistance(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})
	}
}
