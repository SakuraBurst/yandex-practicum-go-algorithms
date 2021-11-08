package goldenMiddle_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	goldenMiddle "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/GoldenMiddle"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestGoldenMiddle(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		goldenMiddle.GoldenMiddle(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}