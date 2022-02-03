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
	tests := []Test{{strings.NewReader(`2 3
101
110
`), "3"}, {strings.NewReader(`3 3
100
110
001
`), "2"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			fieldWithFlowers.FieldWithFlowers(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})
	}
}
