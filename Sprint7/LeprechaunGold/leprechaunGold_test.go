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
	tests := []Test{{strings.NewReader(`5 15
3 8 1 2 5
`), "15"}, {strings.NewReader(`5 19
10 10 7 7 4
`), "18"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			leprechaunGold.LeprechaunGold(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
