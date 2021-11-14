package goldenMiddle_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	goldenMiddle "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/GoldenMiddle"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestGoldenMiddle(t *testing.T) {
	tests := []Test{{strings.NewReader(`2
1
1 3
2
`), "2"}, {strings.NewReader(`2
2
1 2
3 4
`), "2.5"}, {strings.NewReader(`8
10
0 0 0 1 3 3 5 10
4 4 5 7 7 7 8 9 9 10
`), "5"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		goldenMiddle.GoldenMiddle(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
