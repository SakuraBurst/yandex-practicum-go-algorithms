package borderControl_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	borderControl "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/BorderControl"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBorderControl(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		borderControl.BorderControl(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
