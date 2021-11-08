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

func TestBubble(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		bubble.Bubble(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}