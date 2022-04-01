package crib_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	crib "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8FinalEx/Crib"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestCrib(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		crib.Crib(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
