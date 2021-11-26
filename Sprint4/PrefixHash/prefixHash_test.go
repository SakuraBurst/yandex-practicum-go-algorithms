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

func TestPrefixHash(t *testing.T) {
	tests := []Test{{strings.NewReader(`1000
1000009
abcdefgh
7
1 1
1 5
2 3
3 4
4 4
1 8
5 8`), `97
225076
98099
99100
100
436420
193195
`}, {strings.NewReader(`100
10
a
1
1 1`), `7
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		prefixHash.PrefixHash(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
