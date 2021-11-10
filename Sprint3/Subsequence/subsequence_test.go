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

func TestSubsequence(t *testing.T) {
	tests := []Test{{strings.NewReader(`abc
ahbgdcu
`), "True"}, {strings.NewReader(`abcp
ahpc
`), "False"}, {strings.NewReader(`islx
yoytgtshldmogkdburkbcfvoapepjpcuwemusfkfztrzxstytrnarlizjhuoscuzlraezlaweipuuqdgvhwkhhoufexojaps
`), "True"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		subsequence.Subsequence(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
