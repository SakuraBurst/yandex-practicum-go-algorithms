package maximumDistance_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	maximumDistance "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/MaximumDistance"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestMaximumDistance(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		maximumDistance.MaximumDistance(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}