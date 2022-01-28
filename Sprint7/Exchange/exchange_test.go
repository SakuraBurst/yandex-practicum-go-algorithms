package exchange_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	exchange "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/Exchange"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestExchange(t *testing.T) {
	tests := []Test{{strings.NewReader(`6
7 1 5 3 6 4`), "7"}, {strings.NewReader(`5
1 2 3 4 5`), "4"}, {strings.NewReader(`6
1 12 12 16 1 8`), "22"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			exchange.Exchange(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
