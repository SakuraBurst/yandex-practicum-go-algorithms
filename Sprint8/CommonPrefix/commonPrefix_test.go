package commonPrefix_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	commonPrefix "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/CommonPrefix"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestCommonPrefix(t *testing.T) {
	tests := []Test{
		{strings.NewReader(`3
abacaba
abudabi
abcdefg
`), "2"},
		{strings.NewReader(`2
tutu
kukuku
`), "0"},
		{strings.NewReader(`3
qwe
qwerty
qwerpy
`), "3"},
	}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			commonPrefix.CommonPrefix(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})
	}
}
