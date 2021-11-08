package bigInt_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	bigInt "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/BigInt"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBigInt(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		bigInt.BigInt(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}