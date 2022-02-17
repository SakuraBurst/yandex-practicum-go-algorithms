package sameSums_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	sameSums "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7FinalEx/SameSums"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestSameSums(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		sameSums.SameSums(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
