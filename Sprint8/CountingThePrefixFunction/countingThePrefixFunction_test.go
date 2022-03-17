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
	tests := []Test{{strings.NewReader("abracadabrab\n"), "0 0 0 1 0 1 0 1 2 3 4 2 "},
		{strings.NewReader("xxzzxxz\n"), "0 1 0 0 1 2 3 "},
		{strings.NewReader("aaaaa\n"), "0 1 2 3 4 "}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			countingThePrefixFunction.CountingThePrefixFunction(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
