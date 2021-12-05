package filterUp_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	filterUp "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/FilterUp"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestFilterUp(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		filterUp.FilterUp(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}