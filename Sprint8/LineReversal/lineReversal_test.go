package lineReversal_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	lineReversal "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/LineReversal"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestLineReversal(t *testing.T) {
	tests := []Test{{strings.NewReader("one two three"), "three two one"},
		{strings.NewReader("hello"), "hello"},
		{strings.NewReader("may the force be with you"), "you with be force the may"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			lineReversal.LineReversal(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
