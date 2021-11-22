package wierdComparison_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	wierdComparison "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/WierdComparison"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestWierdComparison(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		wierdComparison.WierdComparison(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}