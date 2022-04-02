package packedPrefix_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	packedPrefix "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8FinalEx/PackedPrefix"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestPackedPrefix(t *testing.T) {
	tests := []Test{
		{strings.NewReader(`3
2[a]2[ab]
3[a]2[r2[t]]
a2[aa3[b]]
`), "aaa"},
		{strings.NewReader(`3
abacabaca
2[abac]a
3[aba]
`), "aba"}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		packedPrefix.PackedPrefix(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
