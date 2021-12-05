package numeralPath_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	numeralPath "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/NumeralPath"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestNumeralPath(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		numeralPath.NumeralPath(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}