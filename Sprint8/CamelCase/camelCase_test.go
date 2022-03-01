package camelCase_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	camelCase "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/CamelCase"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestCamelCase(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		camelCase.CamelCase(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
