package substrings_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	substrings "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/Substrings"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestSubstrings(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		substrings.Substrings(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}