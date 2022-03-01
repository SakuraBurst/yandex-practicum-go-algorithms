package countingThePrefixFunction_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	countingThePrefixFunction "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/CountingThePrefixFunction"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestCountingThePrefixFunction(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		countingThePrefixFunction.CountingThePrefixFunction(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
