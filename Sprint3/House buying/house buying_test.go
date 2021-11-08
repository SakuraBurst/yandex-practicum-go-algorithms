package house buying_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	house buying "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/House buying"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestHouse buying(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		house buying.House buying(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}