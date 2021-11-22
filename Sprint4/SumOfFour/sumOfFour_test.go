package sumOfFour_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	sumOfFour "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/SumOfFour"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestSumOfFour(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		sumOfFour.SumOfFour(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}