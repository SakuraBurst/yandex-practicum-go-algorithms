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
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		lineReversal.LineReversal(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
