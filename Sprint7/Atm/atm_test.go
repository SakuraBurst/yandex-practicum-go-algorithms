package atm_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	atm "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/Atm"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestAtm(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		atm.Atm(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
