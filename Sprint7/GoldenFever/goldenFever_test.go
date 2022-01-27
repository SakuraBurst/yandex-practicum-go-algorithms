package goldenFever_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	goldenFever "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/GoldenFever"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestGoldenFever(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		goldenFever.GoldenFever(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
