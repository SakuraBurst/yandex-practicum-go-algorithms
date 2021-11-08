package combinations_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	combinations "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/Combinations"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestCombinations(t *testing.T) {
	tests := []Test{{strings.NewReader("23"), "ad ae af bd be bf cd ce cf"}, {strings.NewReader("92"), "wa wb wc xa xb xc ya yb yc za zb zc"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		combinations.Combinations(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
