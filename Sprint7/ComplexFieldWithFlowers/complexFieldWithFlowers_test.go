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
	tests := []Test{{strings.NewReader(`2 3
101
110
`), `3
URR`}, {strings.NewReader(`3 3
100
110
001
`), `2
UURR`}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			complexFieldWithFlowers.ComplexFieldWithFlowers(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
