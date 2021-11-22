package polinominalHash_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	polinominalHash "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/PolinominalHash"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestPolinominalHash(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		polinominalHash.PolinominalHash(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}