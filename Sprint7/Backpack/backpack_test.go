package backpack_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	backpack "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/Backpack"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBackpack(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		backpack.Backpack(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
