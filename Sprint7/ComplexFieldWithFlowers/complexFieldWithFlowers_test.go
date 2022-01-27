package complexFieldWithFlowers_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	complexFieldWithFlowers "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/ComplexFieldWithFlowers"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestComplexFieldWithFlowers(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		complexFieldWithFlowers.ComplexFieldWithFlowers(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
