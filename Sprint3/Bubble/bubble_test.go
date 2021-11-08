package bubble_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	bubble "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/Bubble"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBubble(t *testing.T) {
	tests := []Test{{strings.NewReader(`5
4 3 9 2 1`), `3 4 2 1 9
3 2 1 4 9
2 1 3 4 9
1 2 3 4 9
`}, {strings.NewReader(`5
12 8 9 10 11`), `8 9 10 11 12
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		bubble.Bubble(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
