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

func TestBipartitionCheck(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		bipartitionCheck.BipartitionCheck(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}