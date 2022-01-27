package leprechaunGold_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	leprechaunGold "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/LeprechaunGold"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestLeprechaunGold(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		leprechaunGold.LeprechaunGold(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
