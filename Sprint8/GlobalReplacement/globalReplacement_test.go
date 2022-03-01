package globalReplacement_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	globalReplacement "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/GlobalReplacement"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestGlobalReplacement(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		globalReplacement.GlobalReplacement(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
