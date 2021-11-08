package subsequence_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	subsequence "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/Subsequence"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestSubsequence(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		subsequence.Subsequence(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}