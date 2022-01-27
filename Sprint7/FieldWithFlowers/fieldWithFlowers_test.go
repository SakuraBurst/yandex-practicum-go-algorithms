package fieldWithFlowers_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	fieldWithFlowers "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/FieldWithFlowers"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestFieldWithFlowers(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		fieldWithFlowers.FieldWithFlowers(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
