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
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		exchange.Exchange(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
