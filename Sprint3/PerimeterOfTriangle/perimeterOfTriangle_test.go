package perimeterOfTriangle_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	perimeterOfTriangle "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/PerimeterOfTriangle"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestPerimeterOfTriangle(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		perimeterOfTriangle.PerimeterOfTriangle(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}