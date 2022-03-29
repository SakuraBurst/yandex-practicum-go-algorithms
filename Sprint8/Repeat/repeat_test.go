package repeat_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	repeat "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/Repeat"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestRepeat(t *testing.T) {
	tests := []Test{
		{strings.NewReader("zzzzzz\n"), "6"},
		{strings.NewReader("abacaba\n"), "1"},
		{strings.NewReader("abababab\n"), "4"},
		{strings.NewReader("tuhhbntuhhbntuhhbn\n"), "3"},
	}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			repeat.Repeat(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
