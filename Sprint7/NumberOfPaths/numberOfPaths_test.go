package numberOfPaths_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	numberOfPaths "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/NumberOfPaths"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestNumberOfPaths(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		numberOfPaths.NumberOfPaths(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
