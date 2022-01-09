package bipartitionCheck_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	bipartitionCheck "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/BipartitionCheck"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBipartitionCheck(t *testing.T) {
	tests := []Test{{strings.NewReader(`3 2
1 2
2 3
`), "YES"}, {strings.NewReader(`3 3
1 2
2 3
1 3
`), "NO"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			bipartitionCheck.BipartitionCheck(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})
	}
}
