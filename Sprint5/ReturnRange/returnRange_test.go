package returnRange_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	returnRange "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/ReturnRange"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestReturnRange(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		returnRange.ReturnRange(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}