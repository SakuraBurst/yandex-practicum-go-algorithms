package prefixHash_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	prefixHash "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/PrefixHash"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestPrefixHash(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		prefixHash.PrefixHash(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}